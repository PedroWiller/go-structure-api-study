package gptService

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PullRequestInc/go-gpt3"

	"gpt-twitter-integration/internal/dto"
	httpRequest "gpt-twitter-integration/pkg/httpRequest"
)

func CreateMessae(message string) ([]dto.Choices, error) {
	apiKey := "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"
	url := "https://api.openai.com/v1/engines/davinci-codex/completions"
	htp := httpRequest.RequestParams{Token: apiKey, Url: url}

	body := dto.Prompt{
		Prompt: message,
		Model:  "gpt-3.5-turbo",
	}

	err, res := htp.Post(body)
	if err != nil {
		return nil, err
	}

	response := res.(dto.ChoicesReponse)

	return response.Choices, nil
}

func TestMessage(prompt string) error {
	apiKey := "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"

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{"The first thing you should know about javascript is"},
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{"."},
		Echo:      true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Choices[0].Text)

	return nil
}

func Execute() {
	apiKey := "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"
	model := "gpt-3.5-turbo" // ou outro modelo disponível

	// Dados da solicitação
	url := "https://api.openai.com/v1/engines/" + model + "/completions"
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + apiKey,
	}
	data := `{
		"prompt": "Olá, GPT-3.5!",
		"max_tokens": 50
	}`

	// Criando a requisição HTTP
	req, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		fmt.Println("Erro ao criar a requisição:", err)
		return
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Enviando a requisição
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao enviar a requisição:", err)
		return
	}
	defer resp.Body.Close()

	// Lendo a resposta
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler a resposta:", err)
		return
	}

	fmt.Println("Resposta:", string(responseBody))
}

