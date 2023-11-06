package weatherRoute

import (
	"github.com/gofiber/fiber/v2"

	weatherController "gpt-twitter-integration/internal/modules/weather/controller"
)

func Init(app *fiber.App) {
	app.Get("/weather", weatherController.Get)
}
