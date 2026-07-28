package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"strconv"
	"time"
)

type Signer struct {
	apiKey string
	secret []byte
}

func NewSigner(apiKey, apiSecret string) Signer {
	return Signer{apiKey, []byte(apiSecret)}
}

func (s Signer) Sign(params url.Values, recvWindowMs int64) string {
	if recvWindowMs > 0 {
		params.Set("recvWindow", strconv.FormatInt(recvWindowMs, 10))
	}
	params.Set("timestamp", strconv.FormatInt(time.Now().UnixMilli(), 10))
	payload := params.Encode()

	mac := hmac.New(sha256.New, s.secret)
	mac.Write([]byte(payload))
	sig := hex.EncodeToString(mac.Sum(nil))

	return payload + "&signature=" + sig
}
