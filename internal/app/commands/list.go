package commands

import (
	"github.com/abozorov/telegramBot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func (c *Commander) List(inputMessage *tgbotapi.Message, productService *product.Service) {
	msgText := "PRODUCT LIST\n\n"
	productList := productService.List()

	for _, p := range productList {
		msgText += p.Title + "\n"
	}
	
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, msgText)
	c.bot.Send(msg)
}