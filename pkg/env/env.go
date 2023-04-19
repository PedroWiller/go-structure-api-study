package env

import "os"

var (
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
	TwitterUrl        string
)

func Init() {
	ConsumerKey = os.Getenv("TWITTER_CONSUMER_KEY")
	ConsumerSecret = os.Getenv("TWITTER_CONSUMER_SECRET")
	AccessToken = os.Getenv("TWITTER_ACCESS_TOKEN")
	AccessTokenSecret = os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	TwitterUrl = os.Getenv("TWITTER_URL")
}
