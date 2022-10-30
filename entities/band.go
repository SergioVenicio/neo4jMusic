package entities

type Band struct {
	Name string
	Musician *Musician
}

func NewBand(name string, musician *Musician) *Band {
	return &Band{
		Name: name,
		Musician: musician,
	}
}