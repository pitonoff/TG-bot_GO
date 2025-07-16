package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные из .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	botToken := os.Getenv("TELEGRAM_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_TOKEN не найден")
	}

	// Инициализация бота
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true // Для логирования

	log.Printf("Бот %s запущен", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID,
					"Привет, "+update.Message.From.FirstName+"! 👋")
				bot.Send(msg)

			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда 😕")
				bot.Send(msg)
			}
		} else {
			// Отвечаем на обычные сообщения
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ты сказал: "+update.Message.Text)
			bot.Send(msg)
		}
	}
}
