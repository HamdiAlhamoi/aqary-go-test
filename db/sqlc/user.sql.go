package db

import (
	"context"
	"database/sql"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    id,
    name,
    phone_number

) VALUES (
    $1,
    $2,
    $3
) 
RETURNING id, name, phone_number, created_at, updated_at;
`

const generateOtp = `-- name: GenerateOtp :one
UPDATE users SET otp = $1, otp_expiration_time = $2
    where phone_number = $3 
	RETURNING id, name, phone_number, otp, otp_expiration_time, created_at, updated_at;

`
const findByPhoneNumber = `-- name: findByPhoneNumber :one
select id from users where phone_number = $1;
`

const validateOtp = `-- name: ValidateOtp :one
SELECT
    *
FROM
    users
WHERE 
phone_number = $1 
 AND otp = $2
    AND otp_expiration_time > NOW();
`

type CreateUserParams struct {
	ID       string `json:"id"`
	Name  string    `json:"name"`
	PhoneNumber     string    `json:"phone_number"`
	
}

type GenerateOtpParams struct {
	PhoneNumber       string `json:"phoneNumber"`
	Otp  string    `json:"otp"`
	OtpExpirationTime       time.Time    `json:"otp_expiration_time"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (*User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.PhoneNumber,

	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.PhoneNumber,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

func (q *Queries) GenerateOtp(ctx context.Context, arg GenerateOtpParams) (*User, error) {
	
	var userID string
	error := q.db.QueryRow(ctx, findByPhoneNumber, arg.PhoneNumber).Scan(&userID)
	if error == sql.ErrNoRows {
		return nil, ErrNotFound
	} else if error != nil {
		return nil, ErrNotFound
	}

	row := q.db.QueryRow(ctx, generateOtp,
		arg.Otp,
		arg.OtpExpirationTime,
		arg.PhoneNumber,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.PhoneNumber,
		&i.Otp,
		&i.OtpExpirationTime,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

func (q *Queries) ValidateOtp(ctx context.Context, arg ValidateOtpParams) (*User, error) {
	row := q.db.QueryRow(ctx, validateOtp,
		arg.PhoneNumber,
		arg.Otp,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.PhoneNumber,
		&i.Otp,
		&i.OtpExpirationTime,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}


type ValidateOtpParams struct {
	PhoneNumber     string    `json:"phone_number"`
	Otp  string    `json:"otp"`
}