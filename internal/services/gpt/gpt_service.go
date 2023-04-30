package gptService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateMessae() {
	apiKey := "YOUR_API_KEY_HERE"
	modelID := "YOUR_MODEL_ID_HERE"
	prompt := "YOUR_PROMPT_HERE"
	headers := http.Header{}
	headers.Add("Content-Type", "application/json")
	headers.Add("Authorization", "Bearer "+apiKey)
	body := struct {
		Prompt string `json:"prompt"`
		Model  string `json:"model"`
	}{Prompt: prompt, Model: modelID}
	requestBody, _ := json.Marshal(body)
	resp, err := http.Post("https://api.openai.com/v1/engines/davinci-codex/completions", "application/json", bytes.NewReader(requestBody))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var response struct {
		Choices []struct {
			Text string `json:"text"`
		} `json:"choices"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		panic(err)
	}
	generatedText := response.Choices[0].Text
	fmt.Println(generatedText)
}
