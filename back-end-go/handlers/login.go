package handlers

import (
	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/nguyenduclam1711/react-signal-chat-app/models"
	"github.com/nguyenduclam1711/react-signal-chat-app/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type channelResult struct {
	result *mongo.SingleResult
	id     string
}

func getUserAndUserCredentialByUsername(username string) (userCredentialResult, userResult *mongo.SingleResult, err error) {
	channel := make(chan channelResult)
	go func() {
		result := repository.UserRepository.GetOne(&bson.D{
			{
				Key:   "username",
				Value: username,
			},
		}, nil)
		channel <- channelResult{
			result: result,
			id:     "UserRepository",
		}
	}()
	go func() {
		result := repository.UserCredentialRepository.GetOne(&bson.D{
			{
				Key:   "username",
				Value: username,
			},
		}, nil)
		channel <- channelResult{
			result: result,
			id:     "UserCredentialRepository",
		}
	}()
	for i := 0; i < 2; i++ {
		currRes := <-channel
		switch currRes.id {
		case "UserRepository":
			userResult = currRes.result
		case "UserCredentialRepository":
			userCredentialResult = currRes.result
		}
	}
	if userResult.Err() != nil {
		err = userResult.Err()
	} else if userCredentialResult.Err() != nil {
		err = userCredentialResult.Err()
	}
	return userCredentialResult, userResult, err
}

func LoginHandler(c *fiber.Ctx) error {
	body := c.Body()

	var parsedPayload models.AuthenticationFrontendPayload
	parseErr := json.Unmarshal(body, &parsedPayload)
	if parseErr != nil {
		return parseErr
	}

	userCredentialResult, userResult, getUserAndCredentialErr := getUserAndUserCredentialByUsername(parsedPayload.Username)
	if getUserAndCredentialErr != nil {
		return getUserAndCredentialErr
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
	return c.JSON(loginResponse)
}
