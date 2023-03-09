package flights

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"smiles/pkg/rabbitmq"
	"smiles/pkg/smiles"
)

type SearchFlightRequest struct {
	Origin      string   `json:"origin"`
	Destination string   `json:"destination"`
	Dates       []string `json:"dates"`
}

type Consumer interface {
	StartConsumer()
}

type FlightConsumer struct {
	Connection *rabbitmq.AmqpClient
}

func (c *FlightConsumer) StartConsumer() {
	args := &rabbitmq.SubscriptionExchange{
		ExchangeName:    "smiles",
		ExchangeType:    "fanout",
		QueueName:       "flights",
		NumberOfWorkers: 1,
	}

	c.Connection.Subscribe(args, worker)
}

func worker(delivery *amqp.Delivery) {
	var search SearchFlightRequest

	if err := json.Unmarshal(delivery.Body, &search); err != nil {
		fmt.Println("Failed to unmarshal request: ", search)
	}

	client := smiles.New()

	client.Search(smiles.PriceRequest{
		Origin:      search.Origin,
		Destination: search.Destination,
		Dates:       search.Dates,
	}).Render()

	delivery.Ack(false)
}

func NewFlightConsumer(client *rabbitmq.AmqpClient) Consumer {
	return &FlightConsumer{
		Connection: client,
	}
}
