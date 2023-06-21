package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/tombuente/tresor/spec/snippetspec"
)

const url = "http://127.0.0.1:8080/api"

var endpoints = map[string]string{
	"snippet": fmt.Sprintf("%s/snippets", url),
}

var _ Tresor = (*TresorImpl)(nil)

type Tresor interface {
	GetSnippet(key string) (snippetspec.SnippetRes, error)
	PostSnippet(body snippetspec.SnippetReq) (snippetspec.SnippetRes, error)
}

type TresorImpl struct{}

func New() TresorImpl {
	return TresorImpl{}
}

func (t TresorImpl) GetSnippet(key string) (snippetspec.SnippetRes, error) {
	res, err := http.Get(fmt.Sprintf("%s/%s", endpoints["snippet"], key))
	if err != nil {
		return snippetspec.SnippetRes{}, errors.New("url error")
	}

	if res.StatusCode != 200 {
		return snippetspec.SnippetRes{}, fmt.Errorf("request status code is not 200, was %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return snippetspec.SnippetRes{}, errors.New("failed to read response body")
	}

	var snippet snippetspec.SnippetRes
	err = json.Unmarshal(body, &snippet)
	if err != nil {
		return snippetspec.SnippetRes{}, errors.New("received malformatted response body")
	}

	return snippet, nil
}

func (t TresorImpl) PostSnippet(snippet snippetspec.SnippetReq) (snippetspec.SnippetRes, error) {
	reqBody, err := json.Marshal(snippet)
	if err != nil {
		return snippetspec.SnippetRes{}, err
	}

	res, err := http.Post(endpoints["snippet"], "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return snippetspec.SnippetRes{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return snippetspec.SnippetRes{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	var snippetRes snippetspec.SnippetRes
	err = json.NewDecoder(res.Body).Decode(&snippetRes)
	if err != nil {
		return snippetspec.SnippetRes{}, err
	}

	return snippetRes, nil
}
