package userController

import (
	"github.com/gofiber/fiber/v2"

	userBusiness "gpt-twitter-integration/internal/modules/user/business"
	"gpt-twitter-integration/pkg/model"
)

func Get(c *fiber.Ctx) error {
	user, err := userBusiness.FindUserByName(c.Query("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Data: user,
	})
}
