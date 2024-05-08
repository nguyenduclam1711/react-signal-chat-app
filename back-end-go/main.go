package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nguyenduclam1711/react-signal-chat-app/database"
	"github.com/nguyenduclam1711/react-signal-chat-app/env"
	"github.com/nguyenduclam1711/react-signal-chat-app/modules"
	"github.com/nguyenduclam1711/react-signal-chat-app/repository"
)

func main() {
	// load env
	env.LoadEnvData()

	// connect to database
	database.ConnectToDatabase()

	// create all repositories
	repository.CreateAllRepositories()

	// create fiber app
	app := fiber.New()

	// craete all modules
	modules.CreateAllModules(app)

	// disconnect database here, the code is ugly I know
	defer func() {
		if disconnectErr := database.MongoDatabase.Client().Disconnect(context.TODO()); disconnectErr != nil {
			log.Fatal("Database disconnect error", disconnectErr)
		}
		database.MongoDatabase = nil
	}()

	log.Fatal(app.Listen(":3000"))
}
