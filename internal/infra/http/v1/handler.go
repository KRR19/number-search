package v1

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/KRR19/number-search/internal/domain/numbersearch"
)

type Handler struct {
	logger       *slog.Logger
	numbersearch *numbersearch.Service
}

func NewHandler(logger *slog.Logger, numbersearch *numbersearch.Service) *Handler {
	return &Handler{
		logger:       logger,
		numbersearch: numbersearch,
	}
}

func (h *Handler) GetNumberPosition(w http.ResponseWriter, r *http.Request) {
	v, err := strconv.Atoi(r.PathValue("number"))
	if err != nil {
		NumberPositionResponse(w, -1, ErrInvalidNumber)
		return
	}

	res, err := h.numbersearch.SearchNumber(r.Context(), v)

	NumberPositionResponse(w, res, err)

}

func NumberPositionResponse(w http.ResponseWriter, position int, err error) {
	status := http.StatusOK
	var er *string

	switch {
	case errors.Is(err, numbersearch.ErrNumberNotFound):
		status = http.StatusNotFound
		errstr := err.Error()
		er = &errstr
	case errors.Is(err, ErrInvalidNumber):
		status = http.StatusBadRequest
		errstr := err.Error()
		er = &errstr
	case err != nil:
		status = http.StatusInternalServerError
		errstr := err.Error()
		er = &errstr
	}

	res := GetNumberPositionResponse{
		Position: position,
		Error:    er,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(res)
}
