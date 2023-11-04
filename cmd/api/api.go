package api

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"gpt-twitter-integration/internal/modules/health"
	"gpt-twitter-integration/internal/server/middleware/routes"
)

func Start() {
	app := fiber.New()
	health.Main(app)

	api := fiber.New()
	routes.Init(api)
	app.Mount("/api/v1", api)

	log.Fatal(app.Listen(":5003"))
}
