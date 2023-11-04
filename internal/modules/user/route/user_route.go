package userRoute

import (
	"github.com/gofiber/fiber/v2"

	userController "gpt-twitter-integration/internal/modules/user/controller"
)

func Init(app *fiber.App) {
	app.Get("/user", userController.Get)
}
