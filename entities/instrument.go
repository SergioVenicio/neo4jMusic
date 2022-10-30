package entities

type Instrument struct {
	Name string
}

func NewInstrument(name string) *Instrument {
	return &Instrument{
		Name: name,
	}
}