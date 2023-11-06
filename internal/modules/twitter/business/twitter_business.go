package twitterBusiness

import (
	"gpt-twitter-integration/internal/modules/twitter/dto"
	twitterService "gpt-twitter-integration/internal/modules/twitter/service"
)

func SendTweet(dto dto.TextTweet) error {
	err := twitterService.SendTweet(dto)
	if err != nil {
		return err
	}

	return nil
}
