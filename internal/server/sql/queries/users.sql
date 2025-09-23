-- name: CreateUser :one
insert into 
    USERS(ID, USERNAME, PASSWORD, CREATED_AT, UPDATED_AT) 
values
    ($1, $2, $3, $4, $5)
returning *;