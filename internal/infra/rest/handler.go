package rest

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
		h.numberPositionResponse(w, -1, ErrInvalidNumber)
		return
	}

	res, err := h.numbersearch.SearchNumber(r.Context(), v)

	h.numberPositionResponse(w, res, err)

}

func (h *Handler) V2GetNumberPosition(w http.ResponseWriter, r *http.Request) {
	v, err := strconv.Atoi(r.PathValue("number"))
	if err != nil {
		h.numberPositionResponse(w, -1, ErrInvalidNumber)
		return
	}

	res, err := h.numbersearch.SearchNumberV2(r.Context(), v)

	h.numberPositionResponse(w, res, err)

}

func (h *Handler) numberPositionResponse(w http.ResponseWriter, position int, err error) {
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
	if err = json.NewEncoder(w).Encode(res); err != nil {
		h.logger.Error("failed to encode response", slog.String("error", err.Error()))
	}
}
