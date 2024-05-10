package modules

import (
	"fmt"
	"sync"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nguyenduclam1711/react-signal-chat-app/env"
	"github.com/nguyenduclam1711/react-signal-chat-app/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateModuleParam struct {
	PrefixPath string
	App        *fiber.App
}

type CreateModuleConstruct struct {
	PrefixPath       string
	CreateModuleFunc func(param CreateModuleParam)
}

func createModules(app *fiber.App, modules []CreateModuleConstruct) {
	if len(modules) < 1 {
		return
	}
	var wg sync.WaitGroup

	wg.Add(len(modules))

	for _, construct := range modules {
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

func createPublicModules(app *fiber.App) {
	publicModules := []CreateModuleConstruct{
		{
			PrefixPath:       "/auth",
			CreateModuleFunc: CreateAuthenticationModule,
		},
		{
			PrefixPath:       "/user",
			CreateModuleFunc: CreatePublicUserModule,
		},
	}
	createModules(app, publicModules)
}

func createAuthModules(app *fiber.App) {
	authModules := []CreateModuleConstruct{
		{
			PrefixPath:       "/chat",
			CreateModuleFunc: CreateChatModule,
		},
		{
			PrefixPath:       "/user",
			CreateModuleFunc: CreateUserModule,
		},
	}
	createModules(app, authModules)
}

func CreateAllModules(app *fiber.App) {
	createPublicModules(app)

	// JWT middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(env.EnvData["JWT_SECRET"]),
		},
		TokenLookup: "cookie:accessToken",
	}))
	// all modules below here's gonna be need bearer token

	createAuthModules(app)
}

func PathWithPrefix(prefix string, endpoint string) string {
	return fmt.Sprint(prefix, endpoint)
}

func CurrentUser(c *fiber.Ctx) models.CurrentUser {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	mongoId, err := primitive.ObjectIDFromHex(claims["id"].(string))
	result := models.CurrentUser{
		Id:       claims["id"].(string),
		Username: claims["username"].(string),
	}
	if err == nil {
		result.MongoId = mongoId
	}
	return result
}
