-- name: CreateArtist :one
INSERT INTO artists (Name)
VALUES ($1)
RETURNING *;

-- name: CreateSongToArtist :one
INSERT INTO artistsToSongs (ArtistID, SongID)
VALUES ($1, $2)
RETURNING *;
