package main

import (
	"gpt-twitter-integration/cmd/api"
	"gpt-twitter-integration/pkg/env"
)

func main() {
	env.Init()
	api.Start()
}
