package routes

import (
	"github.com/gofiber/fiber/v2"

	userRoute "gpt-twitter-integration/internal/modules/user/route"
)

func Init(app *fiber.App) {
	// Users
	userRoute.Init(app)
}
