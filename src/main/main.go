package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/telegram-bot-api.v4"
)

var idAPINiceHash = os.Getenv("ID_API_NICEHASH")
var keyAPINiceHash = os.Getenv("KEY_API_NICEHASH")
var btcAddress = os.Getenv("BTC_ADDRESS")

var urlAPINiceHashBalance = "https://api.nicehash.com/api?method=balance&id=" + idAPINiceHash + "&key=" + keyAPINiceHash
var urlAPINiceHashProvider = "https://api.nicehash.com/api?method=stats.provider&addr=" + btcAddress

var buttons = []tgbotapi.KeyboardButton{
	tgbotapi.KeyboardButton{Text: "Balance"},
	tgbotapi.KeyboardButton{Text: "Provider"},
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

func getProviderNiceHash(url string) string {
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
		case "Provider":
			provider := getProviderNiceHash(urlAPINiceHashProvider)
			log.Printf(provider)
			message = tgbotapi.NewMessage(update.Message.Chat.ID, provider)
		default:
			message = tgbotapi.NewMessage(update.Message.Chat.ID, `test`)
		}

		// В ответном сообщении просим показать клавиатуру
		message.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)

		bot.Send(message)
	}
}
