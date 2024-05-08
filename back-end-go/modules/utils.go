package modules

import (
	"fmt"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nguyenduclam1711/react-signal-chat-app/env"
	"github.com/nguyenduclam1711/react-signal-chat-app/models"
)

type CreateModuleParam struct {
	PrefixPath string
	App        *fiber.App
}

type CreateModuleConstruct struct {
	PrefixPath       string
	CreateModuleFunc func(param CreateModuleParam)
}

func CreateAllModules(app *fiber.App) {
	var wg sync.WaitGroup

	allCreateModuleFuncs := []CreateModuleConstruct{
		{
			PrefixPath:       "/user",
			CreateModuleFunc: CreateUserModule,
		},
		{
			PrefixPath:       "/chat",
			CreateModuleFunc: CreateChatModule,
		},
		{
			PrefixPath:       "/auth",
			CreateModuleFunc: CreateAuthenticationModule,
		},
	}

	// increment the wait group counter
	wg.Add(len(allCreateModuleFuncs))

	for _, construct := range allCreateModuleFuncs {
		// launch each module setup in a go routine
		go func(c CreateModuleConstruct) {
			// defer the counter when the go routine is complete
			defer wg.Done()
			c.CreateModuleFunc(CreateModuleParam{
				PrefixPath: c.PrefixPath,
				App:        app,
			})
		}(construct)
	}

	// wait for all goroutines to complete
	wg.Wait()
}

func PathWithPrefix(prefix string, endpoint string) string {
	return fmt.Sprint(prefix, endpoint)
}

func GenerateAccessToken(user models.UserParseFromDB) (string, error) {
	jwtSecret := []byte(env.EnvData["JWT_SECRET"])
	accessToken := jwt.New(jwt.SigningMethodEdDSA)
	claims := accessToken.Claims.(jwt.MapClaims)
	// 1 week
	claims["exp"] = time.Now().Add(7 * 24 * time.Hour)
	claims["username"] = user.Username
	claims["id"] = user.Id
	return accessToken.SignedString(jwtSecret)
}

func GenerateRefreshToken(accessToken string) (string, error) {
	jwtSecret := []byte(env.EnvData["JWT_SECRET"])
	refreshToken := jwt.New(jwt.SigningMethodEdDSA)
	claims := refreshToken.Claims.(jwt.MapClaims)
	// 1 week + 2 days
	claims["exp"] = time.Now().Add(9 * 24 * time.Hour)
	claims["accessToken"] = accessToken
	return refreshToken.SignedString(jwtSecret)
}

func GenerateAuthTokens(user models.UserParseFromDB) (string, string, error) {
	accessToken, accessTokenErr := GenerateAccessToken(user)
	// FIX: idk why can't generate access token, must investigate into this
	if accessTokenErr != nil {
		return "", "", accessTokenErr
	}
	refreshToken, refreshTokenErr := GenerateRefreshToken(accessToken)
	if refreshTokenErr != nil {
		return accessToken, "", refreshTokenErr
	}
	return accessToken, refreshToken, nil
}
