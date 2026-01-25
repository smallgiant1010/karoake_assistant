-- name: CreateSong :exec
INSERT INTO songs (Language, Title, IsGenerated, Lyrics)
VALUES ($1, $2, $3, $4);

-- name: GetSongByID :one
SELECT * FROM songs
WHERE songs.songID = $1;

-- name: GetSongsByArtist :many
SELECT s.Title, s.Language, s.Lyrics, s.IsGenerated, a.Name
FROM artists a 
LEFT JOIN artistsToSongs j ON a.ArtistID = j.ArtistID
RIGHT JOIN songs s ON j.SongID = s.SongID
WHERE a.Name LIKE $1
ORDER BY a.ArtistID;

-- name: GetSongsByTitle :many
SELECT s.Title, s.Language, s.Lyrics, s.IsGenerated, a.Name
FROM artists a 
LEFT JOIN artistsToSongs j ON a.ArtistID = j.ArtistID
RIGHT JOIN songs s ON j.SongID = s.SongID
WHERE s.Title LIKE $1
ORDER BY s.SongID;


