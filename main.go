package main

import (
	"log"
	"os"

	"fmt"

	"github.com/komisan19/bitbank-checker/bitbank"
	"github.com/line/line-bot-sdk-go/linebot"
)

func floatToString(f []float64) []string {
	i := make([]string, len(f))
	for n := range i {
		i[n] = string(i[n])
	}
	return i
}

func main() {
	bot, err := linebot.New(
	  os.Getenv("LINE_BOT_CHANNEL_SECRET"),
	  os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
	  log.Fatal(err)
	}
	pairs, sums := bitbank.GetTicker()
	for i, sum := range sums {
    result := fmt.Sprintf("%v:%v\n",pairs[i], sum)

	  message := linebot.NewTextMessage(result)
	  if _, err := bot.BroadcastMessage(message).Do(); err != nil {
	    log.Fatal(err)
	  }
	}
}
