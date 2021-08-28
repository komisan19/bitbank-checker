package bitbank

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	json "github.com/goccy/go-json"
)

type Ticker struct {
	Data    []struct {
		Pair      string `json:"pair"`
		Sell      string `json:"sell"`
		Buy       string `json:"buy"`
		Open      string `json:"open"`
		High      string `json:"high"`
		Low       string `json:"low"`
		Last      string `json:"last"`
		Vol       string `json:"vol"`
		Timestamp int64  `json:"timestamp"`
	} `json:"data"`
}

func GetTicker(){
  resp, err := http.Get("https://public.bitbank.cc/tickers")

  if err != nil{
    log.Fatal(err)
  }

  var ticker Ticker

  if err = json.NewDecoder(resp.Body).Decode(&ticker); err != nil{
    fmt.Println(err)
  }
  for _, v := range ticker.Data{
    sell, _ := strconv.ParseFloat(v.Sell, 64)
    open, _ := strconv.ParseFloat(v.Open, 64)
    sum := (sell - open) / open * 100
    fmt.Printf("通貨: %v, 24時間前の始値: %v\n", v.Pair, sum)
  }

  defer resp.Body.Close()

}
