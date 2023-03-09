package views

import (
	"github.com/gofiber/fiber/v2"
	"smiles/internal/scaffold/messagings"
	"smiles/internal/scaffold/views/flights"
)

type AllViews struct {
	Flight flights.View
}

func createViews(producers *messagings.AllMessagingProducers) *AllViews {
	return &AllViews{
		Flight: flights.NewFlightView(producers.Flight),
	}
}

func NewAllViews(app *fiber.App, producers *messagings.AllMessagingProducers) *AllViews {
	views := createViews(producers)

	groupFlights := app.Group("/flights")
	groupFlights.Post("/search", views.Flight.SearchFlightHandler)

	return views
}
