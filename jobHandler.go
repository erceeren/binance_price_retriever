package main

import (
	"github.com/erceeren/binance_price_retriever/binance_client"
	"github.com/erceeren/binance_price_retriever/queue_publisher"
)

func Process(address string, exchange string, ticker_name string) {
	queueHandler := new(queue_publisher.RabbitMQHandler)
	queueHandler.InitHandler(address, exchange)
	ticker := binance_client.GetTicker(ticker_name)
	queueHandler.Publish(ticker)

}
