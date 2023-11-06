package routes

import (
	"github.com/gofiber/fiber/v2"

	userRoute "gpt-twitter-integration/internal/modules/user/route"
	weatherRoute "gpt-twitter-integration/internal/modules/weather/route"
)

func Init(app *fiber.App) {
	// Users
	userRoute.Init(app)

	// Weather
	weatherRoute.Init(app)
}
