package handlers

import (
	"log/slog"

	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/storage"
)

type Handler struct {
	Logger  *slog.Logger
	Storage storage.Queries
}

func NewHandler(logger *slog.Logger, storage storage.Queries) Handler{
	Handler := Handler{
		Logger:  logger,
		Storage: storage,
	}
	return Handler
}