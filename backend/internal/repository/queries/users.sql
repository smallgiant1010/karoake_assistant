-- name: CreateUser :one
INSERT INTO users (Username, Password)
VALUES ($1, $2);

-- name: GetUser :one
SELECT * FROM users
WHERE users.Username == $1;

-- name: IncrementUserCount :one
UPDATE ON users
SET GenerateCount = GenerateCount + 1
WHERE users.UserID == $1;

-- name: CreateUserToSong :one
INSERT INTO usersToSongs (SongID, UserID)
VALUES ($1, $2);

