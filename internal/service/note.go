package service

import (
	"context"

	"github.com/o-ga09/graphql-go/internal/domain"
	"github.com/o-ga09/graphql-go/internal/domain/repository"
	"github.com/o-ga09/graphql-go/internal/service/dto"
)

type NoteService struct {
	noteRepo repository.NoteRepository
}

func NewNoteService(noteRepo repository.NoteRepository) *NoteService {
	return &NoteService{
		noteRepo: noteRepo,
	}
}

func (n *NoteService) FetchNotes(ctx context.Context, userId string) ([]*dto.NoteDto, error) {
	notes, err := n.noteRepo.GetNotes(ctx, userId)
	if err != nil {
		return nil, err
	}

	res := []*dto.NoteDto{}
	for _, note := range notes {
		r := &dto.NoteDto{
			ID:              note.ID,
			UserId:          note.UserID,
			Title:           note.Title,
			Content:         note.Content,
			Tags:            note.Tags,
			CreatedDateTime: note.CreatedDateTime.Format("2006-01-02 15:04:05"),
			UpdatedDateTime: note.UpdatedDateTime.Format("2006-01-02 15:04:05"),
		}
		res = append(res, r)
	}
	return res, nil
}

func (n *NoteService) FetchNoteById(ctx context.Context, id string) (*dto.NoteDto, error) {
	note, err := n.noteRepo.GetNoteByID(ctx, id)
	if err != nil {
		return nil, err
	}
	res := &dto.NoteDto{
		ID:              note.ID,
		UserId:          note.UserID,
		Title:           note.Title,
		Content:         note.Content,
		Tags:            note.Tags,
		CreatedDateTime: note.CreatedDateTime.Format("2006-01-02 15:04:05"),
		UpdatedDateTime: note.UpdatedDateTime.Format("2006-01-02 15:04:05"),
	}
	return res, nil
}

func (n *NoteService) CreateNote(ctx context.Context, note *domain.Note) (*domain.Note, error) {
	err := n.noteRepo.CreateNote(ctx, note)
	if err != nil {
		return nil, err
	}
	cretaedNote, err := n.noteRepo.GetNoteByID(ctx, note.ID)
	if err != nil {
		return nil, err
	}
	return cretaedNote, err
}

func (n *NoteService) UpdateNoteById(ctx context.Context, id string, note *domain.Note) (*domain.Note, error) {
	err := n.noteRepo.UpdateNoteByID(ctx, id, note)
	if err != nil {
		return nil, err
	}
	updatedNote, err := n.noteRepo.GetNoteByID(ctx, id)
	return updatedNote, err
}

func (n *NoteService) DeleteNoteById(ctx context.Context, id string) error {
	err := n.noteRepo.DeleteNoteByID(ctx, id)
	return err
}
