package responces

type ListPhonebooksResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    []Phonebook `json:"data"`
}
type Phonebook struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Options []string `json:"options,omitempty"`
}

// AddPhonebookResponse represents the API response
type AddPhonebookResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
