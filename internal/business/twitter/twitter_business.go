package twitterBusiness

import (
	"gpt-twitter-integration/internal/dto"
	twitterService "gpt-twitter-integration/internal/services/twitter"
)

func CreatePost() error {
	firstChoiceText := "bata choice"
	tweetMessage := dto.PostTextDto{
		Text: firstChoiceText,
	}
	err := twitterService.SendTweet(tweetMessage)
	if err != nil {
		return err
	}

	return nil
}
