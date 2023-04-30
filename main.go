package main

import (
	gptService "gpt-twitter-integration/internal/services/gpt"
	twitterService "gpt-twitter-integration/internal/services/twitter"
	"gpt-twitter-integration/pkg/env"
)

func main() {

	env.Init()

	gptService.Execute()

	err := twitterService.SendTweet()
	if err != nil {
		panic(err)
	}
}
