package ippanel

import (
	"github.com/eqba1/go-smskit/ippanel/responces"
)

// validate the authentication token and get the associated user information.
func (i *Ippanel) CheckToken() (*responces.BaseResponse, error) {
	return i.post("/acl/auth/check_token", nil)
}
