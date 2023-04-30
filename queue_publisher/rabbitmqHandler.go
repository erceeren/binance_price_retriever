package queue_publisher

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/erceeren/binance_price_retriever/binance_client"
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQHandler struct {
	connection *amqp.Connection
	exchange   string
}

func (mq *RabbitMQHandler) Publish(ticker *binance_client.Ticker) error {
	if mq.connection == nil {
		return errors.New("connection is not established")
	}
	channel, _ := mq.connection.Channel()
	content, _ := json.Marshal(ticker)
	channel.PublishWithContext(context.Background(), mq.exchange, "", false, false, amqp.Publishing{Body: content})
	return nil
}

func (mq *RabbitMQHandler) InitHandler(address string, exchange string) error {
	conn, err := amqp.Dial(address)
	if err != nil {
		log.Panicf("couldn't connect to rabbitmq broker: %s", err)
	}
	mq.connection = conn
	mq.exchange = exchange
	return nil
}
