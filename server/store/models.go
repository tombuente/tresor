// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package store

import ()

type CodeLanguage struct {
	ID   int64
	Name string
}

type CodeSnippet struct {
	ID         int64
	Content    string
	LanguageID int64
}