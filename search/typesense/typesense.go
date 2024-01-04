package typesense

import (
	"time"

	ts "github.com/typesense/typesense-go/typesense"

	"github.com/aliml92/realworld-gin-sqlc/config"
)

type TypesenseHandler struct {
	Client     *ts.Client
	Collection string
}

func NewTypesenseHandler(client *ts.Client, collectionName string) *TypesenseHandler {
	return &TypesenseHandler{
		Client:     client,
		Collection: collectionName,
	}
}

func NewClient(config *config.Config) *ts.Client {
	client := ts.NewClient(
		ts.WithConnectionTimeout(5*time.Second),
		ts.WithCircuitBreakerMaxRequests(50),
		ts.WithCircuitBreakerInterval(2*time.Minute),
		ts.WithCircuitBreakerTimeout(1*time.Minute),
	)
	return client
}