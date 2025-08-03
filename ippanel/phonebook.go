package ippanel

import (
	"fmt"

	"github.com/eqba1/go-smskit/ippanel/requests"
	"github.com/eqba1/go-smskit/ippanel/responces"
)

func (i *Ippanel) StorePhonenumber(phoneNumber, pre, fullname, phonebookID string) (*responces.BaseResponse, error) {
	contact := requests.PhonebookContact{
		Number:      phoneNumber,
		Pre:         pre,
		Name:        fullname,
		PhonebookID: phonebookID,
	}

	payload := requests.AddPhonebookRequest{
		List: []requests.PhonebookContact{contact},
	}

	return i.post("/phonebooks/numbers/add-list-new", payload)
}

// CreatePhonebook creates a new phonebook
func (i *Ippanel) CreatePhonebook(title string, options []string) (*responces.BaseResponse, error) {
	if options == nil {
		options = []string{}
	}

	payload := requests.PhonebookRequest{
		Title:   title,
		Options: options,
	}

	return i.post("/phonebooks", payload)
}

func (i *Ippanel) FetchPhonebooks() (*responces.BaseResponse, error) {
	resp, err := i.get("/phonebooks/list-new", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get list phonebooks: %w", err)
	}

	return resp, nil
}
