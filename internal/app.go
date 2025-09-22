package app

import (
	"context"
	"duty_dude/internal/duty_system"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type App struct {
	bot         *bot.Bot
	ctx         context.Context
	cancel      context.CancelFunc
	botName     string
	duty_system duty_system.DutySystem
	log         interface{}
}

func New(token string) (*App, error) {
	ctx, cancel := context.WithCancel(context.Background())

	tgBot, err := bot.New(token, nil)
	if err != nil {
		cancel()
		return nil, err
	}

	botName, err := tgBot.GetMyName(ctx, nil)
	if err != nil {
		cancel()
		return nil, err
	}

	a := App{
		bot:     tgBot,
		botName: botName.Name,
		ctx:     ctx,
		cancel:  cancel,
	}

	a.bot.RegisterHandler(bot.HandlerTypeMessageText, "@"+a.botName, bot.MatchTypeContains, a.onTag, nil)
	return &a, nil
}

func (a *App) onTag(ctx context.Context, b *bot.Bot, update *models.Update) {
	dude, err := a.duty_system.GetDude()
	if err != nil {
		// a.log.Error("Failed to get duty dude", err)
		// a.bot.SendMessage
	}

	a.bot.SendPhoto(a.ctx, &bot.SendPhotoParams{
		ChatID: update.Message.Chat.ID,
		// Photo:   dude.Avatar,
		Caption: dude.Name,
	})
}
