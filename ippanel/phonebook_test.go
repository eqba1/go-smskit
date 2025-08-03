package ippanel

import (
	"fmt"
	"os"
	"testing"
)

// TestSend test send sms
func TestStorePhonenumber(t *testing.T) {
	sms := New(os.Getenv("API_KEY"))

	resp, err := sms.StorePhonenumber("+9810001", "اقای", "امینی", "new")
	if err != nil {
		t.Error("error occurred ", err)
	}

	t.Log(resp)
}

func TestCreatePhonebook(t *testing.T) {
	sms := New(os.Getenv("API_KEY"))

	resp, err := sms.CreatePhonebook("new2", nil)
	if err != nil {
		t.Error("error occurred ", err)
	}

	t.Log(resp)
}

func TestFetchPhonebooks(t *testing.T) {
	sms := New(os.Getenv("API_KEY"))

	resp, err := sms.FetchPhonebooks()
	if err != nil {
		t.Error("error occurred ", err)
	}

	fmt.Println(resp)

	t.Log(resp)
}
