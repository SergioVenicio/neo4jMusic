package entities

type Musician struct {
	Name string
	Instrument *Instrument
}

func NewMusician(name string, instrument *Instrument) *Musician {
	return &Musician{
		Name: name,
		Instrument: instrument,
	}
}