package rest

type SnippetResponse struct {
	Key      string `json:"key"`
	Content  string `json:"content"`
	Language string `json:"language"`
}

type SnippetRequest struct {
	Content  string `json:"content"`
	Language string `json:"language"`
}
