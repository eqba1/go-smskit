package smskit

type SMSKit[T any] interface {
	StorePhonenumber(number, pre, name, title string) (T, error)
	CreatePhonebook(title string, options []string) (T, error)
}
