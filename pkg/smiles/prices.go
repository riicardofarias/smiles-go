package smiles

import (
	"fmt"
)

type PriceRequest struct {
	Origin      string
	Destination string
	Departure   string
}

type BestPrice struct {
	Departure   string
	Origin      string
	Destination string
	Miles       int
}

func GetBestPrices(p PriceRequest) []BestPrice {
	token, err := GetToken()

	if err != nil {
		return nil
	}

	var bestPrices []BestPrice

	flight, err := GetFlights(FlightRequest{
		Origin:      p.Origin,
		Destination: p.Destination,
		Departure:   p.Departure,
		Token:       token,
	})

	if err == nil && len(flight.Flights) > 0 {
		bestPrices = append(bestPrices, BestPrice{
			Departure:   p.Departure,
			Origin:      p.Origin,
			Destination: p.Destination,
			Miles:       flight.Price.Miles,
		})
	}

	for _, airport := range []string{"BSB", "CGH", "GRU", "VCP", "SDU", "CNF", "GIG"} {
		flight, err := GetFlights(FlightRequest{
			Origin:      p.Origin,
			Destination: airport,
			Departure:   p.Departure,
			Token:       token,
		})

		if err != nil || len(flight.Flights) == 0 {
			continue
		}

		flight2, err := GetFlights(FlightRequest{
			Origin:      airport,
			Destination: p.Destination,
			Departure:   p.Departure,
			Token:       token,
		})

		if err != nil || len(flight2.Flights) == 0 {
			continue
		}

		totalMiles := flight.Price.Miles + flight2.Price.Miles

		bestPrices = append(bestPrices, BestPrice{
			Departure:   p.Departure,
			Origin:      fmt.Sprintf("%s > %s", p.Origin, airport),
			Destination: fmt.Sprintf("%s > %s", airport, p.Destination),
			Miles:       totalMiles,
		})
	}

	return bestPrices
}
