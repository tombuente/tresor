package snippetspec

type SnippetRes struct {
	Key      string `json:"key"`
	Content  string `json:"content"`
	Language string `json:"language"`
}

type SnippetReq struct {
	Content  string `json:"content"`
	Language string `json:"language"`
}
