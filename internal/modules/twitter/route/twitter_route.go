package twitterRoute

import (
	"github.com/gofiber/fiber/v2"

	twitterController "gpt-twitter-integration/internal/modules/twitter/controller"
)

func Init(app *fiber.App) {
	app.Post("/tweet", twitterController.Post)
}
