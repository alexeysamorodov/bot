package commands

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	arg, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	product, err := c.ProductService.Get(arg)
	if err != nil {
		log.Println("Not found product by idx:", arg)
		return
	}

	argText := fmt.Sprint("Product:", product)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, argText)

	c.Bot.Send(msg)
}

func init() {
	registeredCommands["get"] = (*Commander).Get
}
