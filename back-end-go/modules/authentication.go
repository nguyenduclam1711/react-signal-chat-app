package modules

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/nguyenduclam1711/react-signal-chat-app/models"
	"github.com/nguyenduclam1711/react-signal-chat-app/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// TODO: for login, logout, etc...
func CreateAuthenticationModule(param CreateModuleParam) {
	param.App.Post(PathWithPrefix(param.PrefixPath, "/login"), func(c *fiber.Ctx) error {
		body := c.Body()

		var parsedPayload models.AuthenticationFrontendPayload
		parseErr := json.Unmarshal(body, &parsedPayload)
		if parseErr != nil {
			return parseErr
		}

		channel := make(chan *mongo.SingleResult)
		go func() {
			channel <- repository.UserRepository.GetOne(&bson.D{
				{
					Key:   "username",
					Value: parsedPayload.Username,
				},
			}, nil)
		}()
		go func() {
			channel <- repository.UserCredentialRepository.GetOne(&bson.D{
				{
					Key:   "username",
					Value: parsedPayload.Username,
				},
			}, nil)
		}()
		userCredentialResult, userResult := <-channel, <-channel
		if userCredentialResult.Err() != nil {
			return userCredentialResult.Err()
		}
		if userResult.Err() != nil {
			return userResult.Err()
		}

		var parsedUserCredential models.UserCredentialParsedFromDB
		parsedUserCredentialErr := userCredentialResult.Decode(&parsedUserCredential)
		if parsedUserCredentialErr != nil {
			return parsedUserCredentialErr
		}

		// compare password and hashed password
		comparePasswordErr := bcrypt.CompareHashAndPassword([]byte(parsedUserCredential.Password), []byte(parsedPayload.Password))
		if comparePasswordErr != nil {
			return comparePasswordErr
		}

		var parsedUser models.UserParseFromDB
		parsedUserErr := userResult.Decode(&parsedUser)
		if parsedUserErr != nil {
			return parsedUserErr
		}

		// generate access token and refresh token
		accessToken, refreshToken, tokensErr := GenerateAuthTokens(parsedUser)
		if tokensErr != nil {
			return tokensErr
		}

		loginResponse := models.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}
		loginResponseStr, encodeLoginResErr := json.Marshal(loginResponse)
		if encodeLoginResErr != nil {
			return encodeLoginResErr
		}
		return c.Send(loginResponseStr)
	})
}
