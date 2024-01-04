package db

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (*User, error)

	GenerateOtp(ctx context.Context, arg GenerateOtpParams) (*User, error)
	ValidateOtp(ctx context.Context, arg ValidateOtpParams) (*User, error)
}

var _ Querier = (*Queries)(nil)
