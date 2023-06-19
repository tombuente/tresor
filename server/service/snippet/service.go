package snippet

import (
	"context"
	"errors"
	"strconv"

	"github.com/tombuente/tresor/store"
)

const encodingBase = 36

type Service struct {
	queries *store.Queries
}

func NewService(queries *store.Queries) Service {
	return Service{queries}
}

func (service Service) GetLanguage(ctx context.Context, id int64) (Language, error) {
	languageDao, err := service.queries.GetCodeLanguage(ctx, id)
	if err != nil {
		return Language{}, errors.New("Language not found")
	}

	return NewLanguage(languageDao.ID, languageDao.Name), nil
}

func (service Service) GetLanguageByName(ctx context.Context, name string) (Language, error) {
	languageDao, err := service.queries.GetCodeLanguageByName(ctx, name)
	if err != nil {
		return Language{}, errors.New("Language not found")
	}

	return NewLanguage(languageDao.ID, languageDao.Name), nil
}

func (service Service) GetSnippet(ctx context.Context, key string) (Snippet, error) {
	id, err := strconv.ParseInt(key, encodingBase, 64)
	if err != nil {
		return Snippet{}, errors.New("bad key")
	}

	snippetDoa, err := service.queries.GetCodeSnippetJoinLanguage(ctx, id)
	if err != nil {
		return Snippet{}, errors.New("Snippet not found")
	}

	return NewSnippet(snippetDoa.ID, strconv.FormatInt(snippetDoa.ID, encodingBase), snippetDoa.Content, NewLanguage(snippetDoa.LangaugeID, snippetDoa.LanguageName)), nil
}

func (service Service) AddSnippet(ctx context.Context, content string, langauge Language) (Snippet, error) {
	snippetDoa, err := service.queries.CreateCodeSnippet(ctx, store.CreateCodeSnippetParams{Content: content, LanguageID: langauge.ID})
	if err != nil {
		return Snippet{}, errors.New("unable to create snippet")
	}

	return NewSnippet(snippetDoa.ID, strconv.FormatInt(snippetDoa.ID, encodingBase), snippetDoa.Content, langauge), nil
}
