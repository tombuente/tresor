// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: queries.sql

package tresor

import (
	"context"
)

const createSnippet = `-- name: CreateSnippet :one
INSERT INTO snippets (content, language)
VALUES (?, ?)
RETURNING id, content, language
`

type CreateSnippetParams struct {
	Content  string
	Language string
}

func (q *Queries) CreateSnippet(ctx context.Context, arg CreateSnippetParams) (SnippetModel, error) {
	row := q.db.QueryRowContext(ctx, createSnippet, arg.Content, arg.Language)
	var i SnippetModel
	err := row.Scan(&i.ID, &i.Content, &i.Language)
	return i, err
}

const getSnippet = `-- name: GetSnippet :one
SELECT id, content, language
FROM snippets
WHERE id = ?
`

func (q *Queries) GetSnippet(ctx context.Context, id int64) (SnippetModel, error) {
	row := q.db.QueryRowContext(ctx, getSnippet, id)
	var i SnippetModel
	err := row.Scan(&i.ID, &i.Content, &i.Language)
	return i, err
}