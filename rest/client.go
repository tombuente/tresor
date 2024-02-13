package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const url = "http://127.0.0.1:8080/api"

var endpoints = map[string]string{
	"snippets": fmt.Sprintf("%s/snippets", url),
}

type Client interface {
	GetSnippet(key string) (SnippetResponse, error)
	PostSnippet(body SnippetRequest) (SnippetResponse, error)
}

var _ Client = (*ClientImpl)(nil)

type ClientImpl struct{}

func New() ClientImpl {
	return ClientImpl{}
}

func (t ClientImpl) GetSnippet(key string) (SnippetResponse, error) {
	res, err := http.Get(fmt.Sprintf("%s/snippets/%s", url, key))
	if err != nil {
		return SnippetResponse{}, errors.New("http get error")
	}

	if res.StatusCode != 200 {
		return SnippetResponse{}, fmt.Errorf("request status code was not 200, was %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return SnippetResponse{}, errors.New("failed to read response body")
	}

	var snippet SnippetResponse
	err = json.Unmarshal(body, &snippet)
	if err != nil {
		return SnippetResponse{}, errors.New("received malformatted response body")
	}

	return snippet, nil
}

func (t ClientImpl) PostSnippet(snippet SnippetRequest) (SnippetResponse, error) {
	reqBody, err := json.Marshal(snippet)
	if err != nil {
		return SnippetResponse{}, err
	}

	res, err := http.Post(endpoints["snippets"], "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return SnippetResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return SnippetResponse{}, fmt.Errorf("request status code was not 200, was %d", res.StatusCode)
	}

	var snippetResponse SnippetResponse
	err = json.NewDecoder(res.Body).Decode(&snippetResponse)
	if err != nil {
		return SnippetResponse{}, err
	}

	return snippetResponse, nil
}
