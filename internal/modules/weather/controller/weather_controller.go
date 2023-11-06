package weatherController

import (
	"github.com/gofiber/fiber/v2"

	weatherBusiness "gpt-twitter-integration/internal/modules/weather/business"
	"gpt-twitter-integration/pkg/model"
)

func Get(c *fiber.Ctx) error {
	weather, err := weatherBusiness.Get(c.Query("name"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Data: weather,
	})
}
