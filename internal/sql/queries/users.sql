-- name: CreateUser :one
insert into 
    USERS(ID, USERNAME, PASSWORD, CREATED_AT, UPDATED_AT) 
values
    ($1, $2, $3, $4, $5)
returning *;

-- name: FindUserById :one
SELECT * FROM users WHERE id = $1;

-- name: FindUserByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: DeleteUserById :execrows
DELETE FROM users WHERE id = $1;

