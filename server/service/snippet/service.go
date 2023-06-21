package snippet

import (
	"context"
	"errors"
	"strconv"

	"github.com/tombuente/tresor/server/store"
)

const encodingBase = 36

var _ Service = (*serviceImpl)(nil)

type Service interface {
	GetLanguage(ctx context.Context, id int64) (Language, error)
	GetLanguageByName(ctx context.Context, name string) (Language, error)

	GetSnippet(ctx context.Context, key string) (Snippet, error)
	AddSnippet(ctx context.Context, content string, langauge Language) (Snippet, error)
}

type serviceImpl struct {
	queries *store.Queries
}

func NewService(queries *store.Queries) serviceImpl {
	return serviceImpl{queries}
}

func (service serviceImpl) GetLanguage(ctx context.Context, id int64) (Language, error) {
	languageDao, err := service.queries.GetCodeLanguage(ctx, id)
	if err != nil {
		return Language{}, errors.New("Language not found")
	}

	return NewLanguage(languageDao.ID, languageDao.Name), nil
}

func (service serviceImpl) GetLanguageByName(ctx context.Context, name string) (Language, error) {
	languageDao, err := service.queries.GetCodeLanguageByName(ctx, name)
	if err != nil {
		return Language{}, errors.New("Language not found")
	}

	return NewLanguage(languageDao.ID, languageDao.Name), nil
}

func (service serviceImpl) GetSnippet(ctx context.Context, key string) (Snippet, error) {
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

func (service serviceImpl) AddSnippet(ctx context.Context, content string, langauge Language) (Snippet, error) {
	snippetDoa, err := service.queries.CreateCodeSnippet(ctx, store.CreateCodeSnippetParams{Content: content, LanguageID: langauge.ID})
	if err != nil {
		return Snippet{}, errors.New("unable to create snippet")
	}

	return NewSnippet(snippetDoa.ID, strconv.FormatInt(snippetDoa.ID, encodingBase), snippetDoa.Content, langauge), nil
}
