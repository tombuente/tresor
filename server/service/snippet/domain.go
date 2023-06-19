package snippet

type Language struct {
	ID   int64  `json:"-"`
	Name string `json:"name"`
}

type Snippet struct {
	ID       int64    `json:"-"`
	Key      string   `json:"key"`
	Content  string   `json:"content"`
	Language Language `json:"language"`
}

func NewLanguage(id int64, name string) Language {
	return Language{
		ID:   id,
		Name: name,
	}
}

func NewSnippet(id int64, key string, content string, language Language) Snippet {
	return Snippet{
		ID:       id,
		Key:      key,
		Content:  content,
		Language: language,
	}
}

// Variant: Instead of using SnippetResponse in the http handler, marshal language as string

// func (s Snippet) MarshalJSON() ([]byte, error) {
// 	type Alias Snippet
// 	return json.Marshal(&struct {
// 		Language string `json:"language"`
// 		Alias
// 	}{
// 		Language: s.Language.Name,
// 		Alias:    (Alias)(s),
// 	})
// }

// func (s Snippet) UnmarshalJSON(data []byte) error {
// 	type Alias Snippet
// 	aux := &struct {
// 		Language string `json:"language"`
// 		Alias
// 	}{
// 		Alias: (Alias)(s),
// 	}

// 	if err := json.Unmarshal(data, &aux); err != nil {
// 		return err
// 	}

// 	s.Language = Language{Name: aux.Language}
// 	return nil
// }
