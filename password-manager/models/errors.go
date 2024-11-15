package models

import "github.com/gofiber/fiber/v2"

const (
	SUCCESS = "SUCCESS"
	FAILED  = "FAILED"
)

type ErrorResponseBody struct {
	Status string       `json:"status"`
	Result interface{}  `json:"result"`
	Error  *fiber.Error `json:"error,omitempty"`
}

func Success(response interface{}) *ErrorResponseBody {
	return &ErrorResponseBody{SUCCESS, response, nil}
}

func Failed(err *fiber.Error) *ErrorResponseBody {
	return &ErrorResponseBody{FAILED, nil, err}
}
