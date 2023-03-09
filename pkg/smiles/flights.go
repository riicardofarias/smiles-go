package smiles

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"log"
	"time"
)

type FlightRequest struct {
	Origin      string
	Destination string
	Date        string
	Token       string
}

type FlightResult struct {
	Segment []Segment `json:"requestedFlightSegmentList"`
}

type Segment struct {
	Type    string   `json:"type"`
	Flights []Flight `json:"flightList"`
	Price   struct {
		Miles int `json:"miles"`
	} `json:"bestPricing"`
}

type Flight struct {
	Departure struct {
		Airport Airport `json:"airport"`
	} `json:"departure"`

	Arrival struct {
		Airport Airport `json:"airport"`
	} `json:"arrival"`
}

type Airport struct {
	Code string `json:"code"`
	Name string `json:"name"`
	City string `json:"city"`
}

func GetFlights(p FlightRequest) (Segment, error) {
	client := resty.New()

	log.Printf("Getting flights from %s to %s on %s", p.Origin, p.Destination, p.Date)

	params := map[string]string{
		"cabin":                  "ECONOMIC",
		"originAirportCode":      p.Origin,
		"destinationAirportCode": p.Destination,
		"departureDate":          p.Date,
		"memberNumber":           "",
		"adults":                 "1",
		"children":               "0",
		"infants":                "0",
		"forceCongener":          "false",
	}

	headers := map[string]string{
		"Accept":    "application/json",
		"x-api-key": "aJqPU7xNHl9qN3NVZnPaJ208aPo2Bh2p2ZV844tw",
	}

	r, err := client.R().
		SetQueryParams(params).
		SetHeaders(headers).
		SetAuthToken(p.Token).
		Get("https://api-air-flightsearch-blue.smiles.com.br/v1/airlines/search")

	time.Sleep(1 * time.Second)

	if err != nil {
		return Segment{}, err
	}

	result := FlightResult{}
	err = json.Unmarshal(r.Body(), &result)

	if err != nil {
		return Segment{}, err
	}

	if len(result.Segment) == 0 {
		return Segment{}, nil
	}

	return result.Segment[0], nil
}
