package botcore

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) ProcessMessage(update tgbotapi.Update) (tgbotapi.MessageConfig, bool) {
	var reply tgbotapi.MessageConfig

	reply = tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	reply.ReplyToMessageID = update.Message.MessageID

	if !update.Message.IsCommand() {
		return reply, false
	}

	cmd := b.cmd.Extract(update.Message.Text)
	log.Println("command:message",
		cmd.Command(),
		":",
		cmd.Message(update.Message.Text),
	)

	return reply, true
}

func (b *Bot) SendMessage(message tgbotapi.MessageConfig) error {
	msg, err := b.core.Send(message)
	if err != nil {
		return err
	}

	log.Printf("Message rec: %s", msg.Text)

	return nil
}
