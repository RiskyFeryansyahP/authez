package model

// ResponseError is a model for custom response error
type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
