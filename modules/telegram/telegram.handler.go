package telegram

import (
	"context"
	"p2ptrader/common"
	"p2ptrader/modules/telegram/dto"
	"slices"
	"strconv"
	"strings"

	"github.com/amarnathcjd/gogram/telegram"
)

type handler struct {
	service Service
}

func newHandler(service Service) common.Handler {
	return handler{service}
}

func (h handler) RegisterHandlers(client *telegram.Client) {
	client.OnCommand("start", h.start)
	client.OnCommand("waitlist", h.waitlist)
	client.OnCommand("connect", h.connect)
}

func (h handler) start(message *telegram.NewMessage) error {
	params := strings.Split(message.Text(), " ")
	if len(params) > 1 && params[1] == "waitlist" {
		if err := h.waitlist(message); err != nil {
			return err
		}
		return nil
	}
	welcomeMessage := `👋 <b>Welcome to Rublex Bot!</b> 

		Tired of constantly juggling tabs, browsers, and apps to manage your P2P trades? We know the struggle. 

		<b>Rublex</b> is the ultimate P2P Exchange Manager. We securely link your exchange accounts via API so you can manage your ads, process orders, and reply to customer chats across multiple platforms—all from right here inside Telegram, powered by AI! 🤖⚡️

		🚧 <b>Project Status: Under Active Development</b>
		We are currently building this project from the ground up, and early support from traders like you gives us the massive push we need to keep maintaining and improving it!

		⏳ <b>Get Early Access!</b>
		Be among the first to try Rublex. Click the command below to join our exclusive waiting list:
		👉 /waitlist

		🔗 <b>Useful Links:</b>
		🌐 <b>Website:</b> <a href="https://rublex.exchange/">rublex.exchange</a>
		💻 <b>GitHub:</b> <a href="https://github.com/aboozaid/rublex-bot">Rublex Repository</a>

		<i>Let’s make P2P trading effortless. We can't wait to show you what we're building!</i> 🚀`
	_, err := message.Reply(welcomeMessage, &telegram.SendOptions{ParseMode: "html"})

	return err
}

func (h handler) waitlist(message *telegram.NewMessage) error {
	err := h.service.RegisterUser(context.Background(), dto.CreateUserParams{
		TelegramID: strconv.FormatInt(message.SenderID(), 10),
		TelegramUsername: &message.Sender.Username,
		FirstName: message.Sender.FirstName,
		LastName: message.Sender.LastName,
		CountryCode: message.Sender.LangCode,
	})
	if err != nil {
		// REVIEW: We should return an error from the service
		if strings.Contains(err.Error(), "must be unique") {
			_, err = message.Reply("You already registered in the waiting list\n\nBest Regards.")
			return err
		}
		return err
	}
	_, err = message.Reply("Thank you for joining our waiting list we appreciate that and looking forward to using our Rublex Bot, We will message you once the bot is ready!\n\nBest Regards.")
	if err != nil {
		return err
	}
	return nil
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
