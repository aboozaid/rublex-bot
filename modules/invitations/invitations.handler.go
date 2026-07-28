package invitations

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"p2ptrader/common"
	"p2ptrader/modules/config"
	"strconv"

	"github.com/amarnathcjd/gogram/telegram"
)

type handler struct {
	service       Service
	configService config.Service
}

func newHandler(service Service, configService config.Service) common.Handler {
	return handler{service, configService}
}

func (h handler) RegisterHandlers(client *telegram.Client) {
	adminID, err := strconv.ParseInt(h.configService.Telegram().TelegramAdminID, 10, 64)
	if err != nil {
		slog.Error("could not parse TELEGRAM_ADMIN_ID value", "error", err)
		os.Exit(1)
	}
	client.OnCommand("invitations_create", h.createInvitation, telegram.FromUsers(adminID))
}

func (h handler) createInvitation(message *telegram.NewMessage) error {
	code, err := h.service.CreateInvitation(context.Background())
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("An invitation has been created\n\nCode: <code>%s</code>\nValid For: 24h", code)
	if _, err := message.Reply(msg); err != nil {
		return err
	}
	return nil
}
