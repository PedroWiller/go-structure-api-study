package worker

import (
	"fmt"

	twitterBusiness "gpt-twitter-integration/internal/business/twitter"
)

func Execute() {
	fmt.Println("Started...")

	err := twitterBusiness.CreatePost()
	if err != nil {
		fmt.Println("Error to execute CreatePost, error details: " + err.Error())
		panic(err)
	}

	fmt.Println("Finished...")
}
