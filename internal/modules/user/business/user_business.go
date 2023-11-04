package userBusiness

import (
	"fmt"

	"gpt-twitter-integration/internal/modules/user/dto"
)

func FindUserByName(id string) (*dto.User, error) {
	if id != "10" {
		return nil, fmt.Errorf("Not found")
	}

	return &dto.User{
		Name: "Pedro Figueredo",
	}, nil
}
