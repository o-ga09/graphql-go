package repository

import (
	"context"

	"github.com/o-ga09/graphql-go/internal/domain"
)

//go:generate moq -out moq/note_repository_mock.go -pkg mock . NoteRepository
type NoteRepository interface {
	GetNoteByUserId(context.Context, string) ([]*domain.Note, error)
	GetNoteAll(context.Context) ([]*domain.Note, error)
	GetNoteByID(ctx context.Context, id string) (*domain.Note, error)
	Save(ctx context.Context, note *domain.Note) error
	Delete(ctx context.Context, id string) error
}
