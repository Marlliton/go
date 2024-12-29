package rabbitmq

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Panicf("erro ao criar conecção %s", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return channel, nil
}

func Consume(ch *amqp.Channel, out chan<- amqp.Delivery) error {
	msgs, err := ch.Consume("minhafila", "go-consume", false, false, false, false, nil)
	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}

	return nil
}

func Publish(ch *amqp.Channel, body string) error {
	err := ch.Publish("amq.direct", "", false, false, amqp.Publishing{
		ContentType: "text/palin",
		Body:        []byte(body),
	})
	fmt.Println("msg publicada")
	if err != nil {
		panic(err)
	}

	return nil
}
