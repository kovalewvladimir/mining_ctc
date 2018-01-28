package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/telegram-bot-api.v4"
)

var buttons = []tgbotapi.KeyboardButton{
	tgbotapi.KeyboardButton{Text: "Balance"},
}

func getBalanceNiceHash(url string) string {
	c := http.Client{}
	resp, err := c.Get(url)
	if err != nil {
		return "API NICEHASH ERROR"
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func main() {
	idAPINiceHash := os.Getenv("ID_API_NICEHASH")
	keyAPINiceHash := os.Getenv("KEY_API_NICEHASH")
	urlAPINiceHashBalance := "https://api.nicehash.com/api?method=balance&id=" + idAPINiceHash + "&key=" + keyAPINiceHash

	const WebhookURL = "https://mining-ctc-bot.herokuapp.com/"
	port := os.Getenv("PORT")
	telegramBotToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(WebhookURL))
	if err != nil {
		log.Fatal(err)
	}

	update := bot.ListenForWebhook("/")
	go http.ListenAndServe(":"+port, nil)

	for update := range update {
		var message tgbotapi.MessageConfig

		log.Println("received text: ", update.Message.Text)

		switch update.Message.Text {
		case "Balance":
			balance := getBalanceNiceHash(urlAPINiceHashBalance)
			log.Printf(balance)
			message = tgbotapi.NewMessage(update.Message.Chat.ID, balance)
		default:
			message = tgbotapi.NewMessage(update.Message.Chat.ID, `test`)
		}

		// В ответном сообщении просим показать клавиатуру
		message.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)

		bot.Send(message)
	}
}
