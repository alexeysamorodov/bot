package commands

import (
	"github.com/alexeysamorodov/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var registeredCommands = map[string]func(c *Commander, msg *tgbotapi.Message){}

type Commander struct {
	Bot            *tgbotapi.BotAPI
	ProductService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		Bot:            bot,
		ProductService: productService,
	}
}

func (c *Commander) HandleUpdate(update *tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	command, ok := registeredCommands[update.Message.Command()]
	if ok {
		command(c, update.Message)
	} else {
		c.Default(update.Message)
	}
}
