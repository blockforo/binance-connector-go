package main

import (
	"context"
	"fmt"

	binance "github.com/blockforo/binance-client"
)

func main() {
	SubAccountList()
}

func SubAccountList() {
	apiKey := "your api key"
	secretKey := "your secret key"
	baseURL := "https://api.binance.com"

	client := binance.NewClient(apiKey, secretKey, baseURL)

	// Query Sub-account List (For Master Account) - /sapi/v1/sub-account/list
	subaccountList, err := client.NewQuerySubAccountListService().Email("test@email.com").
		IsFreeze("").Page(1).Limit(10).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(subaccountList))
}
