package main

import (
	"fmt"
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
	tgbotapi.KeyboardButton{Text: "Выплаченный баланс"},
	tgbotapi.KeyboardButton{Text: "Невыплаченный баланс"},
	tgbotapi.KeyboardButton{Text: "Скорость майнинга"},
}

func main() {
	// TODO: Вынести в переменную окружения
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
		case "Выплаченный баланс":
			var messageText string

			provider := Provider{}
			err := provider.getOfNiceHash(urlAPINiceHashProvider)
			if err != nil {
				log.Print("ERROR provider", err)
				messageText = "**ERROR provider**"
			}

			btcToRub, err := getBTCToRUB()
			if err != nil {
				log.Print("ERROR converter", err)
				messageText = "**ERROR converter**"
			}

			balance := provider.getPaidBalance()

			messageText += "Курс: 1 BTC = **%.2f**\n"
			messageText += "Выплаченный баланс: **%.9f** BTC\n"
			messageText += "Выплаченный баланс: **%.2f** RUB\n"
			messageText = fmt.Sprintf(messageText, btcToRub, balance, balance*btcToRub)

			message = tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
		case "Невыплаченный баланс":
			var messageText string

			provider := Provider{}
			err := provider.getOfNiceHash(urlAPINiceHashProvider)
			if err != nil {
				log.Print("ERROR provider", err)
				messageText = "**ERROR provider**"
			}

			btcToRub, err := getBTCToRUB()
			if err != nil {
				log.Print("ERROR converter", err)
				messageText = "**ERROR converter**"
			}

			balance := provider.getUnpaidBalance()

			messageText += "Курс: 1 BTC = **%.2f**\n"
			messageText += "Невыплаченный баланс: **%.9f** BTC\n"
			messageText += "Невыплаченный баланс: **%.2f** RUB\n"
			messageText = fmt.Sprintf(messageText, btcToRub, balance, balance*btcToRub)

			message = tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
		case "Скорость майнинга":
			var messageText string

			provider := Provider{}
			err := provider.getOfNiceHash(urlAPINiceHashProvider)
			if err != nil {
				log.Print("ERROR provider", err)
				messageText = "**ERROR provider**"
			}

			btcToRub, err := getBTCToRUB()
			if err != nil {
				log.Print("ERROR converter", err)
				messageText = "**ERROR converter**"
			}

			speed := provider.getSpeedMining()

			messageText += "Курс: 1 BTC = **%.2f**\n"
			messageText += "Скорость майнинга: **%.9f** BTC/день\n"
			messageText += "Скорость майнинга: **%.2f** RUB/день\n"
			messageText = fmt.Sprintf(messageText, btcToRub, speed, speed*btcToRub)

			message = tgbotapi.NewMessage(update.Message.Chat.ID, messageText)
		default:
			message = tgbotapi.NewMessage(update.Message.Chat.ID, `**Неверный запрос**`)
		}

		// В ответном сообщении просим показать клавиатуру
		message.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttons)

		bot.Send(message)
	}
}
