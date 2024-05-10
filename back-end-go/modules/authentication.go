package modules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nguyenduclam1711/react-signal-chat-app/handlers"
)

// TODO: for login, logout, etc...
func CreateAuthenticationModule(param CreateModuleParam) {
	param.App.Post(PathWithPrefix(param.PrefixPath, "/login"), func(c *fiber.Ctx) error {
		return handlers.LoginHandler(c)
	})
}
