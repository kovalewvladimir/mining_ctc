package main

import (
	"encoding/json"
	"strconv"
)

var urlAPIConverter = "https://api.cryptonator.com/api/ticker/btc-rub"

type converter struct {
	Ticker struct {
		Base   string `json:"base"`
		Target string `json:"target"`
		Price  string `json:"price"`
		Volume string `json:"volume"`
		Change string `json:"change"`
	} `json:"ticker"`
	Timestamp int    `json:"timestamp"`
	Success   bool   `json:"success"`
	Error     string `json:"error"`
}

func getBTCToRUB() (float64, error) {
	jsonByte, err := getJSONOfURLAPI(urlAPIConverter)
	if err != nil {
		return 0, err
	}

	c := converter{}
	err = json.Unmarshal(jsonByte, &c)
	if err != nil {
		return 0, err
	}

	price, err := strconv.ParseFloat(c.Ticker.Price, 64)
	return price, err
}
