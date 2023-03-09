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

type SearchFlightRequest struct {
	Origin      string   `json:"origin"`
	Destination string   `json:"destination"`
	Date        []string `json:"dates"`
}

func NewFlightView(producer flights.Producer) View {
	return &FlightView{
		producer: producer,
	}
}

func (v *FlightView) SearchFlightHandler(ctx *fiber.Ctx) error {
	var request SearchFlightRequest
	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	message, _ := json.Marshal(request)

	go v.producer.Publish(message)

	return ctx.Status(fiber.StatusOK).JSON(request)
}
