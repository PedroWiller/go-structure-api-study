package gptService

import (
	"gpt-twitter-integration/internal/dto"
	"gpt-twitter-integration/pkg/env"
	httpRequest "gpt-twitter-integration/pkg/httpRequest"
)

func CreateMessae(message string) ([]dto.Choices, error) {
	apiKey := env.OpenAiKey
	url := env.OpenAiUrl
	htp := httpRequest.RequestParams{Token: apiKey, Url: url}

	body := dto.Prompt{
		Prompt: message,
		Model:  env.OpenAiModel,
	}

	err, res := htp.Post(body)
	if err != nil {
		return nil, err
	}

	response := res.(dto.ChoicesReponse)

	return response.Choices, nil
}
