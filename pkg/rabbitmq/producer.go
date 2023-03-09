package rabbitmq

import (
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func (c *AmqpClient) Publish(exchangeName, exchangeType string, message []byte) error {
	if c.Connection == nil {
		panic("Cannot initialize the connection broker. Please initialize the connection first.")
	}

	ch, err := c.Connection.Channel()
	defer ch.Close()

	if err != nil {
		panic("Failed to open a channel: " + err.Error())
	}

	err = ch.ExchangeDeclare(
		exchangeName,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic("Failed to declare an exchange: " + err.Error())
	}

	queue, err := ch.QueueDeclare(
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	err = ch.QueueBind(
		queue.Name,
		"",
		exchangeName,
		false,
		nil,
	)

	err = ch.Publish(
		exchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			Body: message,
		},
	)

	logrus.Infof("A message was sent: %v", string(message))

	return err
}
