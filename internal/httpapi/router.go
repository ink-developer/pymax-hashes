package httpapi

import (
	"pymax-hashes/internal/config"
	"pymax-hashes/internal/storage"
)

type Router struct {
	storage *storage.Storage
	config  *config.Config
}

func New(store *storage.Storage, config *config.Config) *Router {
	return &Router{storage: store, config: config}
}
