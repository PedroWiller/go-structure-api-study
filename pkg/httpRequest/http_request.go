package httprequest

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Request struct {
	Token string `json:"token"`
	Url   string `json:"url"`
}

func (r Request) createHeaders(headers *http.Header) {
	headers.Add("Content-Type", "application/json")
	headers.Add("Authorization", "Bearer "+r.Token)
}

func (r Request) Post(body interface{}, response *interface{}) error {
	headers := http.Header{}
	r.createHeaders(&headers)

	requestBody, _ := json.Marshal(body)
	resp, err := http.Post(r.Url, "application/json", bytes.NewReader(requestBody))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return err
	}

	return nil
}
