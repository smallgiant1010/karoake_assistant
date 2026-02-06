package domains

type Artist struct {
	ArtistID int32
	Name string
}

func NewArtist(artistID_ int32, name_ string) *Artist {
	return &Artist{
		ArtistID: artistID_,
		Name: name_,
	}
}
