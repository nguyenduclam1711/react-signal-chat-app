package modules

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/nguyenduclam1711/react-signal-chat-app/models"
	"github.com/nguyenduclam1711/react-signal-chat-app/repository"
)

func CreateUserModule(param CreateModuleParam) {
	param.App.Post(PathWithPrefix(param.PrefixPath, "/register"), func(c *fiber.Ctx) error {
		requestBody := c.Body()

		var registeredUserPayload models.RegisteredUserPayload
		parseErr := json.Unmarshal(requestBody, &registeredUserPayload)
		if parseErr != nil {
			return parseErr
		}

		// Create user credential
		_, createUserCredentialErr := repository.UserCredentialRepository.InsertOne(registeredUserPayload)
		if createUserCredentialErr != nil {
			return createUserCredentialErr
		}

		// Create user
		createdUser, createdUserErr := repository.UserRepository.InsertOne(models.UserDatabaseStruct{
			Username: registeredUserPayload.Username,
		})
		if createdUserErr != nil {
			return createdUserErr
		}

		// Return the user
		responseJson := models.RegisteredUserResponse{}
		responseJson.Username = registeredUserPayload.Username
		responseJson.Id = models.ConvertObjectIdToString(createdUser.InsertedID)
		encodeResponse, encodingErr := json.Marshal(responseJson)
		if encodingErr != nil {
			return encodingErr
		}
		return c.Send(encodeResponse)
	})
}
