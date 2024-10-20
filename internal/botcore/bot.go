package botcore

import (
	"context"
	"log"
	"tgapi/internal/commands"
	"tgapi/internal/commands/tgcommands"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	core    *tgbotapi.BotAPI
	updates tgbotapi.UpdatesChannel
	ctx     context.Context
	cmd     commands.Commands
}

func Init(token string) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	log.Printf("User info %s", bot.Self.String())

	return &Bot{
		cmd:  commands.New(new(tgcommands.Commands)),
		core: bot,
	}, nil
}

func InitWithContext(ctx context.Context, token string) (*Bot, error) {
	bot, err := Init(token)
	if err != nil {
		return nil, err
	}

	bot.ctx = ctx

	return bot, err
}

func (b *Bot) GetUpdateChan() error {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates, err := b.core.GetUpdatesChan(updateConfig)
	if err != nil {
		return err
	}

	b.updates = updates
	return nil
}

func (b *Bot) CloseUpdateChan() {
	b.updates.Clear()
	b.core.StopReceivingUpdates()
}

func (b *Bot) Updates() <-chan tgbotapi.Update {
	return b.updates
}
