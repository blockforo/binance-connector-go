package main

import (
	"context"
	"fmt"
	"log"

	binance "github.com/blockforo/binance-client"
)

func main() {
	AggregateTrades()
}

func AggregateTrades() {
	client := binance.NewWebsocketAPIClient("", "", "wss://testnet.binance.vision/ws-api/v3")
	err := client.Connect()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	defer client.Close()

	response, err := client.NewAggTradesService().Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	fmt.Println(binance.PrettyPrint(response))

	client.WaitForCloseSignal()
}
