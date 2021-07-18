package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

const PORT = "3000"

func main() {
	port := getEnv("PORT", PORT)
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("[golang] Hello from Heroku!")
	})

	app.Listen(fmt.Sprintf(":%s", port))

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
