package entities

type Music struct {
	Name  string  `json:"name"`
	Order float64 `json:"order"`
}

func NewMusic(name string, order float64) *Music {
	return &Music{
		Name:  name,
		Order: order,
	}
}
