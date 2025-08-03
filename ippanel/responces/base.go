package responces

type BaseResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Meta   Meta        `json:"meta"`
}

type Meta struct {
	Status            bool          `json:"status"`
	Message           string        `json:"message"`
	MessageParameters []interface{} `json:"message_parameters"`
	MessageCode       string        `json:"message_code"`
}
