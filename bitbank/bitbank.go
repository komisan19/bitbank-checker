package bitbank

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	json "github.com/goccy/go-json"
)

type Ticker struct {
	Data []struct {
		Pair string `json:"pair"`
		Sell string `json:"sell"`
		Open string `json:"open"`
	} `json:"data"`
}

func GetTicker() ([]string, []float64) {
	var ticker Ticker
	var resultSum []float64
	var resultPair []string

	resp, err := http.Get("https://public.bitbank.cc/tickers")

	if err != nil {
		log.Fatal(err)
	}

	if err = json.NewDecoder(resp.Body).Decode(&ticker); err != nil {
		fmt.Println(err)
	}

	for _, v := range ticker.Data {
		sell, _ := strconv.ParseFloat(v.Sell, 64)
		open, _ := strconv.ParseFloat(v.Open, 64)
		sum := (sell - open) / open * 100
		resultPair = append(resultPair, v.Pair)
		resultSum = append(resultSum, sum)
  }
	defer resp.Body.Close()

	return resultPair, resultSum
}
