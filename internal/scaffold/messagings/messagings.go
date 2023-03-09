package messagings

import (
	"smiles/internal/scaffold/messagings/flights"
	"smiles/pkg/rabbitmq"
)

type AllMessagingConsumers struct {
	Flight flights.Consumer
}

type AllMessagingProducers struct {
	Flight flights.Producer
}

func StartAllMessagingConsumers(client *rabbitmq.AmqpClient) {
	consumers := AllMessagingConsumers{
		Flight: flights.NewFlightConsumer(client),
	}

	consumers.Flight.StartConsumer()
}

func NewAllMessagingProducers(client *rabbitmq.AmqpClient) *AllMessagingProducers {
	return &AllMessagingProducers{
		Flight: flights.NewFlightProducer(client),
	}
}
