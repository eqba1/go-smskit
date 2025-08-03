package requests

// PhonebookContact represents a contact to be added to phonebook
type PhonebookContact struct {
	Number      string                 `json:"number"`
	Pre         string                 `json:"pre,omitempty"`
	Name        string                 `json:"name"`
	Options     map[string]interface{} `json:"options,omitempty"`
	PhonebookID string                 `json:"phonebook_id"`
}

// AddPhonebookRequest represents the request structure for adding contacts
type AddPhonebookRequest struct {
	List []PhonebookContact `json:"list"`
}

type PhonebookRequest struct {
	ID      int      `json:"id"`
	Title   string   `json:"title"`
	Options []string `json:"options,omitempty"`
}
