package tresor

import (
	"context"
)

type Snippet struct {
	ID       int64
	Content  string
	Language string
}

type Repository interface {
	GetSnippet(ctx context.Context, id int64) (Snippet, error)
	CreateSnippet(ctx context.Context, snippet Snippet) (Snippet, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return Service{
		repository: repository,
	}
}

func (s Service) GetSnippet(ctx context.Context, id int64) (Snippet, error) {
	snippet, err := s.repository.GetSnippet(ctx, id)
	if err != nil {
		return Snippet{}, err
	}

	return snippet, nil
}

func (s Service) CreateSnippet(ctx context.Context, snippet Snippet) (Snippet, error) {
	snippet, err := s.repository.CreateSnippet(ctx, snippet)
	if err != nil {
		return Snippet{}, nil
	}

	return snippet, nil
}
