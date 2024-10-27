package rest

type GetNumberPositionResponse struct {
	Position int     `json:"position"`
	Error    *string `json:"error,omitempty"`
}
