package flights

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"smiles/pkg/rabbitmq"
	"smiles/pkg/smiles"
)

type SearchFlightRequest struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Departure   string `json:"departure"`
}

func (v *SearchFlightRequest) Validate() error {
	if v.Origin == "" {
		return errors.New("origin is required")
	}

	if v.Destination == "" {
		return errors.New("destination is required")
	}

	if v.Departure == "" {
		return errors.New("departure is required")
	}

	return nil
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
		Departure:   search.Departure,
	}).Render()

	delivery.Ack(false)
}

func NewFlightConsumer(client *rabbitmq.AmqpClient) Consumer {
	return &FlightConsumer{
		Connection: client,
	}
}
