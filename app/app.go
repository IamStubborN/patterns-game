package app

import (
	"log"

	"github.com/IamStubborN/patterns-game/bot"
	"github.com/IamStubborN/patterns-game/config"
)

// App - main struct of app
type App struct {
	cfg  config.Config
	tbot *bot.TelegramBot
}

// NewApp - returns new instance of App
func NewApp() *App {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalln(err)
	}
	TBot, err := bot.NewTelegramBot(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	return &App{
		cfg: cfg,
		tbot: TBot,
	}

}

// Run - start the app
func (a *App) Run() {
	a.tbot.CatchUpdates()
}
