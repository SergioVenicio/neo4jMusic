package entities

type Album struct {
	Name string `json:"name"`
	Musics []*Music `json:"musics"`
}

func NewAlbum(name string, musics []*Music) *Album {
	return &Album{
		Name: name,
		Musics: musics,
	}
}