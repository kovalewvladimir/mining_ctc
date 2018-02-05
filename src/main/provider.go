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

// type Provider struct {
// 	resultStructProvider structProvider
// 	PaidBalance          float64
// 	UnpaidBalance        float64
// 	SpeedMining          float64
// }

type Provider structProvider

func (p *Provider) getOfNiceHash(url string) error {
	jsonByte, err := getJSONOfURLAPI(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonByte, p)
	return err
}

//TODO: Пробросить ошибку наружу
func (p Provider) getPaidBalance() float64 {
	var sumBalance float64
	for _, v := range p.Result.Payments {
		balance, err := strconv.ParseFloat(v.Amount, 64)
		if err != nil {
			log.Print("ERROR: ", err)
			return 0
		}
		sumBalance += balance
	}
	return sumBalance
}

//TODO: Пробросить ошибку наружу
func (p Provider) getCommission() float64 {
	var sumCommission float64
	for _, v := range p.Result.Payments {
		commission, err := strconv.ParseFloat(v.Fee, 64)
		if err != nil {
			log.Print("ERROR: ", err)
			return 0
		}
		sumCommission += commission
	}
	return sumCommission
}

//TODO: Пробросить ошибку наружу
func (p Provider) getUnpaidBalance() float64 {
	var sumBalance float64
	for _, v := range p.Result.Stats {
		balance, err := strconv.ParseFloat(v.Balance, 64)
		if err != nil {
			log.Print("ERROR: ", err)
			return 0
		}
		sumBalance += balance
	}
	return sumBalance
}

//TODO: Пробросить ошибку наружу
func (p Provider) getSpeedMining() float64 {
	var sumSpeed float64
	for _, v := range p.Result.Stats {
		speed, err := strconv.ParseFloat(v.AcceptedSpeed, 64)
		if err != nil {
			log.Print("ERROR: ", err)
			return 0
		}
		sumSpeed += speed
	}
	return sumSpeed
}

// func (p Provider) getTelegramMessage() string {
// 	var sumBalance float64
// 	for _, v := range p.Result.Stats {
// 		balance, err := strconv.ParseFloat(v.Balance, 32)
// 		if err != nil {
// 			log.Print("ERROR: ", err)
// 			return "Error: Ошибка при расчете итогового баланса"
// 		}
// 		sumBalance += balance
// 	}
// 	return strconv.FormatFloat(sumBalance, 'g', 1, 64)
// }
