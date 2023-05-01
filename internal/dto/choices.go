package dto

type ChoicesReponse struct {
	Choices []Choices `json:"choices"`
}

type Choices struct {
	Text string `json:"text"`
}
