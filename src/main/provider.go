package main

import (
	"encoding/json"
	"log"
	"strconv"
)

type structProvider struct {
	Result struct {
		Stats []struct {
			Balance       string `json:"balance"`
			RejectedSpeed string `json:"rejected_speed"`
			Algo          int    `json:"algo"`
			AcceptedSpeed string `json:"accepted_speed"`
		} `json:"stats"`
		Payments []struct {
			Amount string `json:"amount"`
			Fee    string `json:"fee"`
			TXID   string `json:"TXID"`
			Time   string `json:"time"`
		} `json:"payments"`
		Addr string `json:"addr"`
	} `json:"result"`
	Method string `json:"method"`
}

type Provider structProvider

func (p *Provider) getOfNiceHash(url string) error {
	jsonByte, err := getJSONOfURLAPI(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonByte, p)
	return err
}

func (p Provider) getTelegramMessage() string {
	var sumBalance float64
	for _, v := range p.Result.Stats {
		balance, err := strconv.ParseFloat(v.Balance, 32)
		if err != nil {
			log.Print("ERROR: ", err)
			return "Error: Ошибка при расчете итогового баланса"
		}
		sumBalance += balance
	}
	return strconv.FormatFloat(sumBalance, 'g', 1, 64)
}
