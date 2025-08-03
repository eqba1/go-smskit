package ippanel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"runtime"

	"github.com/eqba1/go-smskit/ippanel/responces"
)

// get do get request
func (sms Ippanel) get(uri string, params map[string]string) (*responces.BaseResponse, error) {
	return sms.request("GET", uri, params, nil)
}

// post do post request
func (sms Ippanel) post(uri string, data interface{}) (*responces.BaseResponse, error) {
	return sms.request("POST", uri, nil, data)
}

func (i Ippanel) request(method string, uri string, params map[string]string, data interface{}) (*responces.BaseResponse, error) {
	u := i.BaseURL

	// join base url with extra path
	u.Path = path.Join(i.BaseURL.Path, uri)

	// set query params
	p := url.Values{}
	for key, param := range params {
		p.Add(key, param)
	}
	u.RawQuery = p.Encode()

	marshaledBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	requestBody := bytes.NewBuffer(marshaledBody)
	req, err := http.NewRequest(method, u.String(), requestBody)
	if err != nil {
		return nil, err
	}

	//req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", i.APIKey)
	req.Header.Set("User-Agent", "Ippanel/ApiClient/"+ClientVersion+" Go/"+runtime.Version())

	res, err := i.HTTPClient.Do(req)
	if err != nil || res == nil {
		return nil, err
	}

	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusOK, http.StatusCreated:
		_res := &responces.BaseResponse{}
		if err := json.Unmarshal(responseBody, _res); err != nil {
			return nil, fmt.Errorf("could not decode response JSON, %s: %v", string(responseBody), err)
		}

		return _res, nil
	case http.StatusNoContent:
		// Status code 204 is returned for successful DELETE requests. Don't try to
		// unmarshal the body: that would return errors.
		return nil, nil
	case http.StatusInternalServerError:
		// Status code 500 is a server error and means nothing can be done at this
		// point.
		return nil, ErrUnexpectedResponse
	default:
		_res := &responces.BaseResponse{}
		if err := json.Unmarshal(responseBody, _res); err != nil {
			return nil, fmt.Errorf("could not decode response JSON, %s: %v", string(responseBody), err)
		}

		return _res, fmt.Errorf(_res.Meta.Message)
	}
}
