package repository

import (
	"context"

	"github.com/o-ga09/graphql-go/internal/domain"
)

//go:generate moq -out moq/note_repository_mock.go -pkg mock . NoteRepository
type NoteRepository interface {
	GetNotes(context.Context, string) ([]*domain.Note, error)
	GetNoteByID(ctx context.Context, id string) (*domain.Note, error)
	CreateNote(ctx context.Context, note *domain.Note) error
	UpdateNoteByID(ctx context.Context, id string, note *domain.Note) error
	DeleteNoteByID(ctx context.Context, id string) error
}
