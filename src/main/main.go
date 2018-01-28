package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getBalanceNiceHash(url string) string {
	c := http.Client{}
	resp, err := c.Get(url)
	if err != nil {
		return "API NICEHASH ERROR"
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func main() {
	idAPINiceHash := os.Getenv("ID_API_NICEHASH")
	keyAPINiceHash := os.Getenv("KEY_API_NICEHASH")
	urlAPINiceHashBalance := "https://api.nicehash.com/api?method=balance&id=" + idAPINiceHash + "&key=" + keyAPINiceHash

	balance := getBalanceNiceHash(urlAPINiceHashBalance)

	log.Printf(balance)
}
