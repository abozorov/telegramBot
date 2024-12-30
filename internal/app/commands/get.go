package commands

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	arg, err := strconv.Atoi(args)

	if err!= nil {
		log.Println("WRONG ARGS", args)
		return
	}

	prod, err := c.productService.Get(arg)

	if err != nil {
		log.Printf("fail to get prod with id %d %v", arg, err)
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		prod.Title,
	)

	c.bot.Send(msg)
}