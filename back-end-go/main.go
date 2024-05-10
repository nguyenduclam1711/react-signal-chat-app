package main

import (
	"context"
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

	// create fiber app with custom json encoder and decoder
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Initialize default config
	app.Use(compress.New(compress.Config{
		Level: compress.LevelDefault,
	}))

	// logger
	app.Use(logger.New(logger.ConfigDefault))

	// create all modules
	modules.CreateAllModules(app)

	// disconnect database here, the code is ugly I know
	defer func() {
		if disconnectErr := database.MongoDatabase.Client().Disconnect(context.TODO()); disconnectErr != nil {
			log.Fatal("Database disconnect error", disconnectErr)
		}
		database.MongoDatabase = nil
	}()

	log.Fatal(app.Listen(":" + env.EnvData["PORT"]))
}
