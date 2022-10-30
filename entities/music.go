package entities

type Music struct {
	Name string `json:"name"`
	Order int64 `json:"order"`
}

func NewMusic(name string, order int64) *Music {
	return &Music{
		Name: name,
		Order: order,
	}
}