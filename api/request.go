package api

import (
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *RestClient) General() Query[Response] {
	return NewQuery[Response](c).
		WithMethod(http.MethodPost).
		WithPath("/")
}
