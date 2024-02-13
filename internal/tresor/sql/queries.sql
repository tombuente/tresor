-- name: GetSnippet :one
SELECT *
FROM snippets
WHERE id = ?;

-- name: CreateSnippet :one
INSERT INTO snippets (content, language)
VALUES (?, ?)
RETURNING *;
