package smiles

import (
	"fmt"
)

type PriceRequest struct {
	Origin      string
	Destination string
	Dates       []string
}

type BestPrice struct {
	Date        string
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

	for _, date := range p.Dates {
		flight, err := GetFlights(FlightRequest{
			Origin:      p.Origin,
			Destination: p.Destination,
			Date:        date,
			Token:       token,
		})

		if err == nil && len(flight.Flights) > 0 {
			bestPrices = append(bestPrices, BestPrice{
				Date:        date,
				Origin:      p.Origin,
				Destination: p.Destination,
				Miles:       flight.Price.Miles,
			})
		}

		for _, airport := range []string{"BSB", "CGH", "GRU", "VCP", "SDU", "CNF", "GIG"} {
			flight, err := GetFlights(FlightRequest{
				Origin:      p.Origin,
				Destination: airport,
				Date:        date,
				Token:       token,
			})

			if err != nil || len(flight.Flights) == 0 {
				continue
			}

			flight2, err := GetFlights(FlightRequest{
				Origin:      airport,
				Destination: p.Destination,
				Date:        date,
				Token:       token,
			})

			if err != nil || len(flight2.Flights) == 0 {
				continue
			}

			totalMiles := flight.Price.Miles + flight2.Price.Miles

			bestPrices = append(bestPrices, BestPrice{
				Date:        date,
				Origin:      fmt.Sprintf("%s > %s", p.Origin, airport),
				Destination: fmt.Sprintf("%s > %s", airport, p.Destination),
				Miles:       totalMiles,
			})
		}
	}

	return bestPrices
}
