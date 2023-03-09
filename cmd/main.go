package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"smiles/config"
	"smiles/internal/scaffold/messagings"
	"smiles/internal/scaffold/views"
	"smiles/pkg/rabbitmq"
)

func main() {
	config.InitConfigs()

	cfg := rabbitmq.GetRabbitMQConfig()

	broker := rabbitmq.New()
	broker.ConnectToBroker(cfg)

	app := fiber.New()
	app.Use(recover.New())

	allProducers := messagings.NewAllMessagingProducers(broker)
	messagings.StartAllMessagingConsumers(broker)

	views.NewAllViews(app, allProducers)

	app.Listen(":8080")
}
