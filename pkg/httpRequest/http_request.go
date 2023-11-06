package httpRequest

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type RequestParams struct {
	Token   string
	Url     string
	Headers *http.Header
}

func (r RequestParams) AddHeaders(headers *http.Header) {
	headers.Add("Content-Type", "application/json")
	headers.Add("Authorization", "Bearer "+r.Token)
}

func (r RequestParams) Post(body interface{}) (error, any) {
	headers := http.Header{}
	r.AddHeaders(&headers)

	requestBody, _ := json.Marshal(body)
	resp, err := http.Post(r.Url, "application/json", bytes.NewReader(requestBody))
	if err != nil {
		return err, nil
	}

	defer resp.Body.Close()

	var T any
	err = json.NewDecoder(resp.Body).Decode(&T)
	if err != nil {
		return err, nil
	}

	return nil, T
}

func (r RequestParams) Get(response interface{}) error {
	resp, err := http.Get(r.Url)
	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil
	}

	return nil
}
