package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BinanceResponse struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
}

func GetFromBinance(identifier string) (float64, error) {
	// Set up the request URL
	apiURL := "https://api.binance.com/api/v3/ticker/price"
	url := fmt.Sprintf("%s?symbol=%s", apiURL, identifier)

	// Create an HTTP client
	client := &http.Client{}

	// Create a GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	// Read the response body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	// Parse the response JSON
	var BinanceResponse BinanceResponse
	err = json.Unmarshal(body, &BinanceResponse)
	if err != nil {
		return 0, err
	}

	return BinanceResponse.Price, nil
}
