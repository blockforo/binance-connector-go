package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	AccountApiTradingStatus()
}

func AccountApiTradingStatus() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// AccountApiTradingStatusService - /sapi/v1/account/apiTradingStatus
	accountApiTradingStatus, err := client.NewAccountApiTradingStatusService().
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(accountApiTradingStatus))
}
