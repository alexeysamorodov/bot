package commands

import (
	"github.com/alexeysamorodov/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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
