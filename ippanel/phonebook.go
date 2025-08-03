package ippanel

import (
	"errors"
	"fmt"

	"github.com/eqba1/go-smskit/ippanel/requests"
	"github.com/eqba1/go-smskit/ippanel/responces"
)

func (i *Ippanel) StorePhonenumber(number, pre, name, title string) (*responces.BaseResponse, error) {
	// Check if phonebook exist get id or create new phonebook and get id
	phonebookID, err := i.getPhonebookID(title)
	if err != nil {
		if errors.Is(err, ErrPhonebookNotFound) {
			phonebookID, err = i.createNewPhonebook(title)
			if err != nil {
				return nil, err
			}
		}
		return nil, err
	}

	// Create the contact object
	contact := requests.PhonebookContact{
		Number:      number,
		Pre:         pre,
		Name:        name,        // Assuming user has Name field
		PhonebookID: phonebookID, // Assuming user has PhonebookID field
	}

	// Create the request payload
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

	payload := requests.CreatePhonebookRequest{
		Title:   title,
		Options: options,
	}

	return i.post("/phonebooks", payload)
}

func (i *Ippanel) getPhonebookID(title string) (id string, err error) {
	phonebooks, err := i.listPhonebooks()
	if err != nil {
		return "", fmt.Errorf("failed to list phonebooks: %w", err)
	}

	for _, phonebook := range phonebooks {
		if phonebook.Title == title {
			return phonebook.ID, nil
		}
	}

	return "", ErrPhonebookNotFound
}

func (i *Ippanel) createNewPhonebook(title string) (id string, err error) {
	requestData := requests.CreatePhonebookRequest{
		Title:   title,
		Options: nil,
	}

	resp, err := i.post("/phonebooks", requestData)
	if err != nil {
		return "", fmt.Errorf("failed to create phonebook: %w", err)
	}

	if !resp.Meta.Status {
		return "", fmt.Errorf("API returned error: %s", resp.Meta.Message)
	}

	// Extract ID from response data
	dataMap, ok := resp.Data.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("unexpected response data format")
	}

	id, ok = dataMap["id"].(string)
	if !ok {
		return "", fmt.Errorf("phonebook ID not found in response")
	}

	return id, nil
}

func (i *Ippanel) listPhonebooks() ([]responces.Phonebook, error) {
	resp, err := i.get("/phonebooks/list-new", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get list phonebooks: %w", err)
	}

	if !resp.Meta.Status {
		return nil, fmt.Errorf("API returned error: %s", resp.Meta.Message)
	}

	// Convert response data to phonebooks slice
	dataSlice, ok := resp.Data.([]interface{})
	if !ok {
		return nil, fmt.Errorf("unexpected response data format")
	}

	phonebooks := make([]responces.Phonebook, len(dataSlice))
	for i, item := range dataSlice {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("unexpected phonebook item format")
		}

		phonebook := responces.Phonebook{}
		if id, exists := itemMap["id"]; exists {
			phonebook.ID = fmt.Sprintf("%v", id)
		}
		if title, exists := itemMap["title"]; exists {
			phonebook.Title = fmt.Sprintf("%v", title)
		}
		if options, exists := itemMap["options"]; exists {
			if optSlice, ok := options.([]interface{}); ok {
				phonebook.Options = make([]string, len(optSlice))
				for j, opt := range optSlice {
					phonebook.Options[j] = fmt.Sprintf("%v", opt)
				}
			}
		}

		phonebooks[i] = phonebook
	}

	return phonebooks, nil
}
