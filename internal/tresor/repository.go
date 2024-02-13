package tresor

import (
	"context"
	"errors"

	"github.com/jinzhu/copier"
)

var _ Repository = (*dbImpl)(nil)

type dbImpl struct {
	queries *Queries
}

func NewDB(queries *Queries) dbImpl {
	return dbImpl{
		queries: queries,
	}
}

func (db dbImpl) GetSnippet(ctx context.Context, id int64) (Snippet, error) {
	dbSnippet, err := db.queries.GetSnippet(ctx, id)
	if err != nil {
		return Snippet{}, errors.New("unable to get code snippet from database")
	}

	var snippet Snippet
	copier.Copy(&snippet, &dbSnippet)
	return snippet, nil
}

func (db dbImpl) CreateSnippet(ctx context.Context, snippet Snippet) (Snippet, error) {
	var params CreateSnippetParams
	copier.Copy(&params, &snippet)
	dbSnippet, err := db.queries.CreateSnippet(ctx, params)
	if err != nil {
		return Snippet{}, err
	}

	var newSnippet Snippet
	copier.Copy(&newSnippet, dbSnippet)
	return newSnippet, nil
}
