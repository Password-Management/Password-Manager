package models

import "github.com/gofiber/fiber/v2"

const (
	SUCCESS = "SUCCESS"
	FAILED  = "FAILED"
)

type ResponseBody struct {
	Status string       `json:"status"`
	Result interface{}  `json:"result"`
	Error  *fiber.Error `json:"error,omitempty"`
}

func Success(response interface{}) *ResponseBody {
	return &ResponseBody{SUCCESS, response, nil}
}

func Failed(err *fiber.Error) *ResponseBody {
	return &ResponseBody{FAILED, nil, err}
}
