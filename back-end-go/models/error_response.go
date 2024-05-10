package models

import (
	"github.com/goccy/go-json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewErrorResponse() ErrorResponse {
	return ErrorResponse{}
}

func GetErrorJsonResponse(message string, status int) []byte {
	errorResponse := NewErrorResponse()
	errorResponse.Message = message
	errorResponse.Status = status
	result, err := json.Marshal(errorResponse)
	if err != nil {
		str := fmt.Sprintf("{\"status\": %v, \"message\": %v}", status, message)
		return []byte(str)
	}
	return result
}

func GetBadRequestErrorJsonResponse() []byte {
	return GetErrorJsonResponse("Bad Request", fiber.StatusBadRequest)
}

func GetUnProcessableEntityJsonResponse() []byte {
	return GetErrorJsonResponse("Unprocessable Entity", fiber.StatusUnprocessableEntity)
}
