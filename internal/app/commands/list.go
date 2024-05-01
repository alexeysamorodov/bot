package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsg := "Products: \n\n"
	products := c.ProductService.List()

	for _, p := range products {
		outputMsg += p.Title + "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsg)

	c.Bot.Send(msg)
}
