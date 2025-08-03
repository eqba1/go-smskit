package smskit

type SMSKit[T any] interface {
	StorePhonenumber(phonenumber, pre, fullname, phonebookID string) (T, error)
	CreatePhonebook(title string, options []string) (T, error)
	FetchPhonebooks() (T, error)
	CheckToken() (T, error)
}
