-- name: CreateArtist :exec
INSERT INTO artists (Name)
VALUES ($1);

-- name: CreateSongToArtist :exec
INSERT INTO artistsToSongs (ArtistID, SongID)
VALUES ($1, $2);
