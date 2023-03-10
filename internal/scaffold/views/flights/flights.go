package flights

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"smiles/internal/scaffold/messagings/flights"
)

type View interface {
	SearchFlightHandler(ctx *fiber.Ctx) error
}

type FlightView struct {
	producer flights.Producer
}

func NewFlightView(producer flights.Producer) View {
	return &FlightView{
		producer: producer,
	}
}

func (v *FlightView) SearchFlightHandler(ctx *fiber.Ctx) error {
	var request flights.SearchFlightRequest
	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	if err := request.Validate(); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	message, _ := json.Marshal(request)

	go v.producer.Publish(message)

	return ctx.Status(fiber.StatusOK).JSON(request)
}
