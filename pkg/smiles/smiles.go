package smiles

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

type Smiles struct {
	Results []BestPrice
}

func New() *Smiles {
	return &Smiles{}
}

func (s *Smiles) Search(p PriceRequest) *Smiles {
	s.Results = GetBestPrices(p)
	return s
}

func (s *Smiles) GetResults() []BestPrice {
	return s.Results
}

func (s *Smiles) Render() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Date", "Origin", "Destination", "Miles"})

	for _, flight := range s.Results {
		t.AppendRow(table.Row{flight.Departure, flight.Origin, flight.Destination, flight.Miles})
	}

	t.Render()
}
