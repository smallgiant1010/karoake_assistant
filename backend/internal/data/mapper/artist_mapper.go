package mapper

import (
	"karoake_assistant/backend/internal/domains"
	"karoake_assistant/backend/internal/data/sqlc"
)

func ArtistModelToDomain(model *sqlc.Artist) *domains.Artist {
	return domains.NewArtist(
		model.Artistid,
		model.Name,
	)
}

func ArtistDomainToModel(domain *domains.Artist) *sqlc.Artist {
	return &sqlc.Artist{
		Artistid: domain.ArtistID,
		Name: domain.Name,
	}
}
