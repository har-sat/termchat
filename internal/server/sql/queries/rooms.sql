-- name: CreateRoom :one
INSERT INTO rooms(id, name, created_at, creator_id, owner_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteRoom :exec
DELETE FROM rooms where id = $1;

-- name: FindAllRooms :many
SELECT * FROM rooms;

-- name: FindRoomById :one
SELECT * FROM rooms WHERE id = $1;