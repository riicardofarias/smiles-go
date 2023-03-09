package rabbitmq

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type SubscriptionExchange struct {
	ExchangeName    string
	ExchangeType    string
	QueueName       string
	NumberOfWorkers int
}

func (c *AmqpClient) Subscribe(args *SubscriptionExchange, handler func(delivery *amqp.Delivery)) {
	args.Validate()

	if c.Connection == nil {
		panic("Cannot initialize the connection broker. Please initialize the connection first.")
	}

	ch, err := c.Connection.Channel()
	if err != nil {
		panic("Failed to open a channel: " + err.Error())
	}

	if err = ch.ExchangeDeclare(
		args.ExchangeName,
		args.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		panic("Failed to declare an exchange: " + err.Error())
	}

	queue, err := ch.QueueDeclare(
		args.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic("Failed to declare a queue: " + err.Error())
	}

	if err = ch.QueueBind(
		queue.Name,
		"",
		args.ExchangeName,
		false,
		nil,
	); err != nil {
		panic("Failed to bind a queue to an exchange: " + err.Error())
	}

	if err = ch.Qos(
		1,
		0,
		false,
	); err != nil {
		panic("Failed to set QoS: " + err.Error())
	}

	deliveries, err := ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	for i := 1; i <= args.NumberOfWorkers; i++ {
		go Worker(deliveries, handler)
	}

	logrus.Infof("Workers started: %d", args.NumberOfWorkers)
}

func Worker(delivery <-chan amqp.Delivery, handler func(delivery *amqp.Delivery)) {
	for d := range delivery {
		handler(&d)
	}
}

func (c SubscriptionExchange) Validate() {
	if c.ExchangeName == "" {
		panic("Exchange name cannot be empty")
	}

	if c.ExchangeType == "" {
		panic("Exchange type cannot be empty")
	}

	if c.QueueName == "" {
		panic("Queue name cannot be empty")
	}

	if c.NumberOfWorkers <= 0 {
		panic("Number of workers must be greater than 0")
	}
}
