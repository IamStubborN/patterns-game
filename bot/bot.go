package bot

import (
	"log"

	"github.com/IamStubborN/patterns-game/config"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
)

//TelegramBot - telegram bot struct
type TelegramBot struct {
	Bot *tg.BotAPI
}

//NewTelegramBot - return telegram bot instance
func NewTelegramBot(cfg config.Config) (*TelegramBot, error) {
	bot, err := tg.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		return nil, err
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &TelegramBot{Bot: bot}, nil
}

//CatchUpdates - Run catch updates from telegram
func (tb *TelegramBot) CatchUpdates() error {
	u := tg.NewUpdate(0)
	u.Timeout = 60

	updates, err := tb.Bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tg.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		tb.Bot.Send(msg)
	}
	return nil
}
