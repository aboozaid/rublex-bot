package binance

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"p2ptrader/modules/binance/dto"
	"p2ptrader/modules/http"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/imroc/req/v3"
)

type strategy struct {
	offset atomic.Int64
}

func newStrategy() http.HttpStrategy {
	return &strategy{}
}

func (s *strategy) ParseHeaders(resp *req.Response) (int, int) {
	if resp.Request == nil || resp.Request.URL == nil {
		return 0, 0
	}
	path := resp.Request.URL.Path
	var limit int
	var usedStr string

	if strings.HasPrefix(path, "/sapi/") {
		limit = 12000

		// Binance sometimes sends sapi-ip-weight, but safely fallback to mbx-used-weight
		usedStr = resp.GetHeader("x-sapi-used-ip-weight-1m")
		if usedStr == "" {
			usedStr = resp.GetHeader("x-mbx-used-weight-1m")
		}
	} else {
		limit = 6000
		usedStr = resp.GetHeader("x-mbx-used-weight-1m")
	}

	used, err := strconv.Atoi(usedStr)
	if err != nil {
		return 0, limit
	}

	return used, limit
}

func (s *strategy) ShouldRetry(resp *req.Response, err error) bool {
	if err != nil {
		return true
	}

	if resp.GetStatusCode() == 429 || resp.GetStatusCode() >= 500 {
		return true
	}

	if resp.GetStatusCode() == 418 {
		// TODO: We should not reach this but for now skip retrying entirly
		return false
	}

	var bodyErr dto.BinanceErrorResponse
	if err := resp.UnmarshalJson(&bodyErr); err == nil {
		switch bodyErr.Code {
		case -1003, -1021, -1022:
			if err := s.syncTime(resp.Request.GetClient()); err != nil {
				slog.Error("failed to sync time on retrying", "error", err)
				return false
			}
			time.Sleep(3 * time.Second)
			return true
		case -9000:
			pattern := `\b(187040|187039|187031)\b`
			re := regexp.MustCompile(pattern)
			matches := re.FindAllString(bodyErr.Message, -1)
			if len(matches) > 0 {
				return false
			}
			// TODO: Log this to debug later
			return true
		case -1000:
			// unknown error
			return true
		default:
			// TODO: Log this to debug later
			return false
		}
	}
	return false
}

func (s *strategy) OnBeforeRequest(req *req.Request) error {
	keys, ok := getBinanceKeys(req.Context())
	if !ok {
		return nil
	}

	values := req.URL.Query()
	values.Set("recvWindow", "60000")

	timeOffset := s.offset.Load()
	if timeOffset == 0 {
		if err := s.syncTime(req.GetClient()); err != nil {
			return err
		}
		timeOffset = s.offset.Load()
	}
	localNow := time.Now().UnixMilli()
	adjustedTimestamp := localNow + timeOffset

	values.Set("timestamp", fmt.Sprintf("%d", adjustedTimestamp))

	payload := values.Encode()
	signature := createSignature(payload, keys.ApiSecret)

	query := payload + "&signature=" + signature
	req.URL.RawQuery = query

	req.Headers.Set("clientType", "web")
	req.Headers.Set("X-MBX-APIKEY", keys.ApiKey)

	return nil
}

func (s *strategy) IsRequestSuccess(resp *req.Response, err error) bool {
	statusCode := resp.GetStatusCode()
	if err == nil || statusCode < 300 {
		return true
	}
	// TODO: We should log to debug for any usecases
	return false
}

func (s *strategy) syncTime(client *req.Client) error {
	// Record local time BEFORE the network call
	localTimeBefore := time.Now().UnixMilli()

	resp, err := client.R().Get("/api/v3/time")
	if err != nil {
		return err
	}

	var result struct {
		ServerTime int64 `json:"serverTime"`
	}
	if err := json.Unmarshal(resp.Bytes(), &result); err != nil {
		return err
	}

	// Calculate network latency (Round Trip Time)
	rtt := time.Now().UnixMilli() - localTimeBefore

	// Estimated Binance time when the response arrived
	// (Server time + half of the network trip time)
	adjustedBinanceTime := result.ServerTime + (rtt / 2)

	// Calculate the exact offset: How much do we need to add to our local clock?
	offset := adjustedBinanceTime - time.Now().UnixMilli()

	// Safely store it
	s.offset.Store(offset)

	fmt.Printf("Binance time synced. Offset applied: %dms\n", offset)
	return nil
}

/*func (s *strategy) OnAppStart() error {
	runCtx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	go func() {
		ticker := time.NewTicker(s.syncDuration)
		defer ticker.Stop()
		defer close(s.stopped)

		for {
			select {
			case <-runCtx.Done():
				return
			case <-ticker.C:
				if err := s.syncTime(runCtx); err != nil {
					slog.Error("failed to sync time of binance server", "error", err)
					return
				}
			}
		}
	}()
	return nil
}

func (s *strategy) OnAppStop() error {
	if s.cancel != nil {
		s.cancel()
	}

	<-s.stopped
	return nil
}*/
