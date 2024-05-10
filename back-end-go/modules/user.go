package modules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nguyenduclam1711/react-signal-chat-app/models"
	"github.com/nguyenduclam1711/react-signal-chat-app/repository"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUserModule(param CreateModuleParam) {
	param.App.Get(PathWithPrefix(param.PrefixPath, "/me"), func(c *fiber.Ctx) error {
		currentUser := CurrentUser(c)
		result := repository.UserRepository.GetOne(&bson.D{{
			Key:   "_id",
			Value: currentUser.MongoId,
		}}, nil)

		var userFromDb models.UserParseFromDB
		decodeErr := result.Decode(&userFromDb)
		if decodeErr != nil {
			return decodeErr
		}
		return c.JSON(userFromDb)
	})
}
