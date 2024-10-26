package v1

import "net/http"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetNumberPosition(w http.ResponseWriter, r *http.Request) {
	v := r.PathValue("number")

	w.Write([]byte(v))
}
