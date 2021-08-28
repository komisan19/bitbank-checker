package main

import (
	"log"
	"os"

	"github.com/komisan19/bitbank-checker/bitbank"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
  bot, err := linebot.New(
    os.Getenv("LINE_BOT_CHANNEL_SECRET"),
    os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
  )
  if err != nil {
    log.Fatal(err)
  }
  result := bitbank.GetTicker()
  message := linebot.NewTextMessage(result)
  if _, err := bot.BroadcastMessage(message).Do(); err != nil {
    log.Fatal(err)
  }
}
