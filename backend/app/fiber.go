package app

import "github.com/gofiber/fiber/v3"

func CreateFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024, // 10 MB
	})
	return app
}
