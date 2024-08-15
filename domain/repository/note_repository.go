package repository

import "github.com/o-ga09/graphql-go/domain"

type NoteRepository interface {
	GetNotes() ([]*domain.Note, error)
	GetNoteByID(id string) (*domain.Note, error)
	CreateNote(note *domain.Note) error
	UpdateNoteByID(id string, note *domain.Note) error
	DeleteNoteByID(id string) error
}
