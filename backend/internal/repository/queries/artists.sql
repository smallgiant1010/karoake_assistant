-- name: CreateArtist :one
INSERT INTO artists (Name)
VALUES ($1);

-- name: CreateSongToArtist :one
INSERT INTO artistsToSongs (ArtistID, SongID)
VALUES ($1, $2);
