package db

import (
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Name  string    `json:"name"`
	PhoneNumber     string    `json:"phone_number"`
	Otp  string    `json:"otp"`
	OtpExpirationTime       time.Time    `json:"otp_expiration_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
