-- name: GetCodeSnippet :one
SELECT *
FROM code_snippets
WHERE id = ?;

-- name: GetCodeSnippetJoinLanguage :one
SELECT code_snippets.id, code_snippets.content, code_languages.id AS langauge_id, code_languages.name AS language_name
FROM code_snippets
JOIN code_languages
ON code_snippets.language_id = code_languages.id
WHERE code_snippets.id = ?;


-- name: CreateCodeSnippet :one
INSERT INTO code_snippets (
    content, language_id
) VALUES (
    ?, ?
)
RETURNING *;

-- name: GetCodeLanguage :one
SELECT * FROM code_languages
WHERE id = ?;

-- name: GetCodeLanguageByName :one
SELECT * FROM code_languages
WHERE name = ?;
