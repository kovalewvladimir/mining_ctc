package main

import (
	"io/ioutil"
	"net/http"
)

func getJSONOfURLAPI(url string) ([]byte, error) {
	c := http.Client{}
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
