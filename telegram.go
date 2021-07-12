package main

import (
	"bytes"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/kbinani/screenshot"
	"image/png"
	"log"
)

func startTelegramClient(config *AppConfig) {
	//Создаем бота
	bot, err := tgbotapi.NewBotAPI(config.BotId)
	if err != nil {
		panic(err)
	}

	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		switch update.Message.Text {
		case "/screen":
			img, _ := screenshot.CaptureDisplay(0)
			var buf bytes.Buffer
			if err := png.Encode(&buf, img); err != nil {
				log.Println("Screenshot failed", err)
			}
			var body = tgbotapi.FileBytes{Name: "screen.png", Bytes: buf.Bytes()}
			msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, body)
			bot.Send(msg)
		case "/gold":
			for _, v := range config.GoldInBags {
				body := tgbotapi.FileBytes{Name: v.Name + ".png", Bytes: gold[v.Name]}
				msg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, body)
				msg.Caption = v.Name
				bot.Send(msg)
			}
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "loh pidr, не пиши сюда больше!")
			bot.Send(msg)
		}
	}
}
