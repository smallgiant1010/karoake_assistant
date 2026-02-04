-- name: CreateUser :one
INSERT INTO users (Username, Password)
VALUES ($1, $2)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE users.Username = $1;

-- name: IncrementUserCount :exec
UPDATE users
SET GenerateCount = GenerateCount + 1
WHERE users.UserID = $1;

-- name: CreateUserToSong :one
INSERT INTO usersToSongs (SongID, UserID)
VALUES ($1, $2)
RETURNING *;

