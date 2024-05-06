package modules

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/nguyenduclam1711/react-signal-chat-app/models"
	"github.com/nguyenduclam1711/react-signal-chat-app/repository"
)

func CreateUserCredentialModule(app *fiber.App) {
	app.Post("/user-credential", func(c *fiber.Ctx) error {
		bodyRequest := c.Body()
		userCredentialFrontendPayload := models.NewUserCredentialFrontendPayload()
		err := json.Unmarshal(bodyRequest, &userCredentialFrontendPayload)
		if err != nil {
			errorResponse := models.NewErrorResponse()
			errorResponse.Message = "Bad request"
			errorResponse.Status = fiber.StatusBadRequest
			errorPayload, errEncodePayload := json.Marshal(errorResponse)

			if errEncodePayload != nil {
				return errEncodePayload
			}
			return c.Status(errorResponse.Status).Send(errorPayload)
		}
		result, createErr := repository.UserCredentialRepository.Create(userCredentialFrontendPayload)
		if createErr != nil {
			return createErr
		}
		createdUserCredentialResponse, errCreatedUserCredentialRes := json.Marshal(result)
		if errCreatedUserCredentialRes != nil {
			return errCreatedUserCredentialRes
		}
		return c.Send(createdUserCredentialResponse)
	})
}
