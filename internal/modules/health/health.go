package health

import (
	"github.com/gofiber/fiber/v2"

	"gpt-twitter-integration/pkg/model"
)

func Main(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
			Data: "OK",
		})
	})
}
