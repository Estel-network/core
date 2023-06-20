package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CoinGeckoResponse struct {
	MarketData struct {
		CurrentPrice map[string]float64 `json:"current_price"`
	} `json:"market_data"`
}

func GetFromGecko(identifier string) (float64, error) {
	// Make the API request
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s", identifier)
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	// Parse the JSON response
	var cgResponse CoinGeckoResponse
	err = json.Unmarshal(body, &cgResponse)
	if err != nil {
		return 0, err
	}

	// Extract the price from the response
	price := cgResponse.MarketData.CurrentPrice["usd"]

	return price, nil
}
