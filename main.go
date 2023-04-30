package main

import "os"

func main() {
	address := os.Getenv("rabbitmq_address")
	exchange := os.Getenv("exchange_name")
	ticker_name := os.Getenv("ticker_name")
	Process(address, exchange, ticker_name)
}
