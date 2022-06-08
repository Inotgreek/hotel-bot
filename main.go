package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type personData struct {
	State  int //0-name,1-number,2-email,3-room,4-auto
	Name   string
	Number string
	Email  string
	Room   string
	Auto   string
}

var personDataMap map[int]*personData

func init() {
	personDataMap = make(map[int]*personData)
}

var mainMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Услуги"),
		tgbotapi.NewKeyboardButton("Ресторан"),
		tgbotapi.NewKeyboardButton("Трансфер"),
		tgbotapi.NewKeyboardButton("SOS"),
	),
)

var serviceMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Уборка"),
		tgbotapi.NewKeyboardButton("Будильник"),
		tgbotapi.NewKeyboardButton("Стирка"),
		tgbotapi.NewKeyboardButton("Кроватка1"),
	),
)

var serviceMenuCleaningTime = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("8-10"),
		tgbotapi.NewKeyboardButton("10-12"),
		tgbotapi.NewKeyboardButton("12-14"),
		tgbotapi.NewKeyboardButton("14-16"),
	),
)

var transferKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Краснодар", "На какое время?"),
		tgbotapi.NewInlineKeyboardButtonData("Аэропорт", "Аэропорт"),
		tgbotapi.NewInlineKeyboardButtonData("Вокзал", "Вокзал"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5441073477:AAFk65WEjUjgVOITvSWgl7hMHL31IlfbXY0")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}
	for update := range updates {

		if update.CallbackQuery != nil {

			bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data))

			queryMSG := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)

			bot.Send(queryMSG)
		}

		if update.Message != nil {

			if update.Message.IsCommand() {
				cmdText := update.Message.Command()

				if cmdText == "start" {
					personDataMap[update.Message.From.ID] = new(personData)
					personDataMap[update.Message.From.ID].State = 0
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Регистрация. Введите имя.")
					bot.Send(msg)
				} else {
				}

				if cmdText == "menu" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Главное меню")
					msg.ReplyMarkup = mainMenu
					bot.Send(msg)
				}
			} else {
				command := update.Message.Text

				switch command {
				case "Трансфер":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Куда хотите поехать?")
					msg.ReplyMarkup = transferKeyboard
					bot.Send(msg)
				case "Услуги":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите услугу")
					msg.ReplyMarkup = serviceMenu
					bot.Send(msg)
				case "Уборка":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите время")
					msg.ReplyMarkup = serviceMenuCleaningTime
					bot.Send(msg)
				default:
					pd, ok := personDataMap[update.Message.From.ID]
					if ok {
						if pd.State == 0 {
							pd.Name = update.Message.Text
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите телефон")
							bot.Send(msg)
							pd.State = 1
						} else if pd.State == 1 {
							pd.Number = update.Message.Text
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите email")
							bot.Send(msg)
							pd.State = 2
						} else if pd.State == 2 {
							pd.Email = update.Message.Text
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите номер комнаты")
							bot.Send(msg)
							pd.State = 3
						} else if pd.State == 3 {
							pd.Number = update.Message.Text
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите номер авто")
							bot.Send(msg)
							pd.State = 4
						} else if pd.State == 4 {
							pd.Auto = update.Message.Text
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Регистрация завершена. Введите /menu.")
							bot.Send(msg)
							pd.State = 5
						}
					}

				}

			}
		}
	}
}
