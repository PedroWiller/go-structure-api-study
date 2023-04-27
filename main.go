package main

import (
	"fmt"

	twitterService "gpt-twitter-integration/internal/services"
	"gpt-twitter-integration/pkg/env"
)

func main() {
	fmt.Println("Say! hello!")

	env.Init()

	err := twitterService.SendTweet()
	if err != nil {
		panic(err)
	}
}
