package main

import (
	"fmt"
	"log"
	"os"

	// https://github.com/maxim-avramenko/monitoring
	// https://github.com/go-telegram-bot-api/telegram-bot-api
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// http poll 是客户端向telegram 服务器定时不断的请求，使用简单，一般用于测试，在本地机器即可完成。老外们的经验说，这种模式经常卡，导致业务处理延迟。
// webhook 是telegram 服务器主动发消息到自己指定的网址，一般用于真实环境，而且对外需要HTTPS / TLS 协议，用户延迟小、并发高。
func main() {
	// for key, value := range os.Environ() {
	// 	fmt.Println(key, value)
	// }
	// fmt.Println(os.Getenv("TELEGRAM_APITOKEN"))
	// return
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))

	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		fmt.Println("%v", update)
		log.Printf("ChatID %s : [%s] %s", update.Message.Chat.ID, update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text+"计算中。。。")
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
