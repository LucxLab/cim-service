package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
)

const defaultHost = "localhost:5672"
const defaultCredentials = "guest:guest"

type Publisher struct {
	Channel *amqp091.Channel
}

func NewPublisher() *Publisher {
	conn, err := amqp091.Dial("amqp://" + defaultCredentials + "@" + defaultHost)
	if err != nil {
		panic(err) //TODO: handle error
	}

	channel, err := conn.Channel()
	if err != nil {
		panic(err) //TODO: handle error
	}
	return &Publisher{
		Channel: channel,
	}
}
