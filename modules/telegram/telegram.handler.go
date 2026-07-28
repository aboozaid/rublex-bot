package telegram

import (
	"context"
	"p2ptrader/common"
	"p2ptrader/modules/telegram/dto"
	"slices"
	"strconv"

	"github.com/amarnathcjd/gogram/telegram"
)

type handler struct {
	service Service
}

func newHandler(service Service) common.Handler {
	return handler{service}
}

func (h handler) RegisterHandlers(client *telegram.Client) {
	client.OnCommand("connect", h.connect)
}

func (h handler) connect(message *telegram.NewMessage) error {
	tgGroupID := strconv.FormatInt(message.ChatID(), 10)
	exists, err := h.service.IsGroupAlreadyUsed(context.Background(), tgGroupID)
	if err != nil {
		return err
	}
	if exists {
		_, err := message.Reply("This group is already connected")
		return err
	}
	member, err := message.Client.GetChatMember(message.ChatID(), message.SenderID())
	if err != nil {
		_, err := message.Reply("Unable to connect the bot, ensure you added the bot as admin before connecting")
		return err
	}
	if member.Status != "creator" {
		_, err := message.Reply("Only group owner can connect the bot")
		return err
	}

	botMember, _ := message.Client.GetChatMember(message.ChatID(), message.Client.Me().ID)
	rights := []bool{
		botMember.Rights.DeleteMessages,
		botMember.Rights.ManageTopics,
		//botMember.Rights.PostMessages,
		//botMember.Rights.EditMessages,
		botMember.Rights.ChangeInfo,
		botMember.Rights.PinMessages,
		//botMember.Rights.ManageDirectMessages,
	}
	if botMember.Status != "admin" || slices.Contains(rights, false) {
		_, err := message.Reply("Please give the bot all required permissions in order to connect\n\nHere are rights that the bot need:\n1. Change group info\n2. Delete messages\n3. Pin messages\n4. Edit member tags\n5. Manage topics (enable group forum in order to see it)")

		return err
	}

	conv, err := message.Client.NewConversation(message.ChatID())
	if err != nil {
		return err
	}
	defer conv.Close()

	tgID := strconv.FormatInt(message.SenderID(), 10)
	accounts, err := h.service.GetUserAccounts(context.Background(), tgID)
	if err != nil {
		return err
	}

	choices := make([]string, len(accounts))
	for i, a := range accounts {
		choices[i] = a.Account
	}
	callback, err := conv.Choice("Select an account to link the group for", choices)
	if err != nil {
		return err
	}
	selectedChoice := string(callback.Data)
	choiceIndex := slices.Index(choices, selectedChoice)

	createGroupParams := dto.CreateGroupParams{GroupID: message.ChatID(), Title: message.Chat.Title}
	if err := h.service.CreateGroup(context.Background(), accounts[choiceIndex].ID, createGroupParams); err != nil {
		return err
	}

	_, err = conv.Reply("The group connected successfully to your account")

	return err
}
