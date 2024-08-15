package repository

import (
	"context"

	"github.com/o-ga09/graphql-go/domain"
)

type NoteRepository interface {
	GetNotes(context.Context) ([]*domain.Note, error)
	GetNoteByID(ctx context.Context, id string) (*domain.Note, error)
	CreateNote(ctx context.Context, note *domain.Note) error
	UpdateNoteByID(ctx context.Context, id string, note *domain.Note) error
	DeleteNoteByID(ctx context.Context, id string) error
}
