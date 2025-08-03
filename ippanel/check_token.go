package ippanel

import (
	"github.com/eqba1/go-smskit/ippanel/responces"
)

func (i *Ippanel) CheckToken() (*responces.BaseResponse, error) {
	// validate the authentication token and get the associated user information.
	return i.post("/acl/auth/check_token", nil)
}
