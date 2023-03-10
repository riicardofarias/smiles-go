package flights

import (
	"github.com/sirupsen/logrus"
	"smiles/pkg/rabbitmq"
)

type Producer interface {
	Publish(request []byte) error
}

type FlightProducer struct {
	Connection *rabbitmq.AmqpClient
}

func NewFlightProducer(client *rabbitmq.AmqpClient) Producer {
	return &FlightProducer{
		Connection: client,
	}
}

func (p *FlightProducer) Publish(request []byte) error {
	err := p.Connection.Publish("smiles", "fanout", "flights", request)
	if err != nil {
		logrus.Errorf("Failed to publish message: %s", string(request))
		return err
	}

	return nil
}
