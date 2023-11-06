package twitterController

import (
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"

	twitterBusiness "gpt-twitter-integration/internal/modules/twitter/business"
	"gpt-twitter-integration/internal/modules/twitter/dto"
	"gpt-twitter-integration/pkg/model"
)

func Post(c *fiber.Ctx) error {
	var dto dto.TextTweet
	if err := json.Unmarshal(c.Body(), &dto); err != nil {
		return errors.New("invalid json")
	}
	err := twitterBusiness.SendTweet(dto)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Data: dto,
	})
}
