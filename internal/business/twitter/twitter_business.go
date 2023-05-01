package twitterBusiness

import (
	"gpt-twitter-integration/internal/dto"
	gptService "gpt-twitter-integration/internal/services/gpt"
	twitterService "gpt-twitter-integration/internal/services/twitter"
)

func CreatePost() error {
	messageContext := "The first sith from history"
	twitterBusiness.CreatePost()
	choices, err := gptService.CreateMessae(messageContext)
	if err != nil {
		return err
	}

	firstChoiceText := choices[0].Text
	tweetMessage := dto.PostTextDto{
		Text: firstChoiceText,
	}
	err = twitterService.SendTweet(tweetMessage)
	if err != nil {
		return err
	}

	return nil
}
