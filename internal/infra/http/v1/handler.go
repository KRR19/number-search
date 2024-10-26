package v1

import (
	"log/slog"
	"net/http"
)

type Handler struct {
	logger *slog.Logger
}

func NewHandler(logger *slog.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

func (h *Handler) GetNumberPosition(w http.ResponseWriter, r *http.Request) {
	v := r.PathValue("number")

	if _, err := w.Write([]byte(v)); err != nil {
		h.logger.Error("error writing response", "error", err)
	}
}
