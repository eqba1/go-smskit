package responces

import "encoding/json"

// BaseResponse base response model
type BaseResponse struct {
	Data json.RawMessage     `json:"data"`
	Meta *PaginationResponse `json:"meta"`
}

type PaginationResponse struct {
	CurrentPage       int           `json:"current_page"`
	From              int           `json:"from"`
	LastPage          int           `json:"last_page"`
	PerPage           int           `json:"per_page"`
	To                int           `json:"to"`
	Total             int           `json:"total"`
	Status            bool          `json:"status"`
	Message           string        `json:"message"`
	MessageParameters []interface{} `json:"message_parameters"`
	MessageCode       string        `json:"message_code"`
}
