package twitterService

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/dghubble/oauth1"
	json "github.com/dustin/gojson"

	"gpt-twitter-integration/internal/modules/twitter/dto"
	"gpt-twitter-integration/pkg/env"
)

func authenticate() *http.Client {
	config := oauth1.NewConfig(env.ConsumerKey, env.ConsumerSecret)
	token := oauth1.NewToken(env.AccessToken, env.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	resp, err := httpClient.Get(fmt.Sprintf("%s/2/users/1234567890", env.TwitterUrl))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return httpClient
}

func SendTweet(dto dto.TextTweet) error {
	values := url.Values{}
	values.Set("tweet_mode", "extended")

	client := authenticate()

	json, _ := json.Marshal(dto)

	buf := bytes.NewBuffer(json)

	resp, err := client.Post(fmt.Sprintf("%s/2/tweets", env.TwitterUrl), "application/json", buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
