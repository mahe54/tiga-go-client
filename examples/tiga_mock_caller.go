package main

import (
	"errors"
	"net/http"
)

type MockCaller struct {
	BaseURL string
}

func (cc *MockCaller) DoCall(req *http.Request) (*http.Response, error) {

	return nil, errors.New("case not found - ?")
}
