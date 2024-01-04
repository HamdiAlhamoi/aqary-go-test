-- name: CreateUser :one
INSERT INTO users (
    id,
    name,
    phone_number,
    otp,
    otp_expiration_time, 
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) 
RETURNING *;