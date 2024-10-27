package v1

type GetNumberPositionResponse struct {
	Position int     `json:"position"`
	Error    *string `json:"error,omitempty"`
}
