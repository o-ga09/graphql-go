package service

import (
	"context"

	"github.com/o-ga09/graphql-go/domain"
	"github.com/o-ga09/graphql-go/domain/repository"
)

type NoteService struct {
	noteRepo repository.NoteRepository
}

func NewNoteService(noteRepo repository.NoteRepository) *NoteService {
	return &NoteService{
		noteRepo: noteRepo,
	}
}

func (n *NoteService) FetchNotes(ctx context.Context) ([]*domain.Note, error) {
	notes, err := n.noteRepo.GetNotes(ctx)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (n *NoteService) FetchNoteById(ctx context.Context, id string) (*domain.Note, error) {
	note, err := n.noteRepo.GetNoteByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (n *NoteService) CreateNote(ctx context.Context, note *domain.Note) error {
	err := n.noteRepo.CreateNote(ctx, note)
	return err
}

func (n *NoteService) UpdateNoteById(ctx context.Context, id string, note *domain.Note) error {
	err := n.noteRepo.UpdateNoteByID(ctx, id, note)
	return err
}

func (n *NoteService) DeleteNoteById(ctx context.Context, id string) error {
	err := n.noteRepo.DeleteNoteByID(ctx, id)
	return err
}
