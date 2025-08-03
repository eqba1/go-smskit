package ippanel

import "fmt"

var (
	ErrPhonebookNotFound  = fmt.Errorf("phonebook not found")
	ErrUnexpectedResponse = fmt.Errorf("inputs have some problems")
)
