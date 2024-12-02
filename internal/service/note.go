package service

import (
	"context"
	"log/slog"
	"strings"

	"github.com/o-ga09/graphql-go/internal/domain"
	"github.com/o-ga09/graphql-go/internal/domain/repository"
	"github.com/o-ga09/graphql-go/internal/service/dto"
	"github.com/o-ga09/graphql-go/pkg/logger"
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
			CreatedDateTime: note.CreatedDateTime,
			UpdatedDateTime: note.UpdatedDateTime,
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
		CreatedDateTime: note.CreatedDateTime,
		UpdatedDateTime: note.UpdatedDateTime,
	}
	return res, nil
}

func (n *NoteService) CreateNote(ctx context.Context, note *domain.Note) (*dto.NoteDto, error) {
	err := n.noteRepo.Save(ctx, note)
	if err != nil {
		return nil, err
	}

	createdNote, err := n.noteRepo.GetNoteByID(ctx, note.ID)
	if err != nil {
		return nil, err
	}

	slog.Log(ctx, logger.SeverityInfo, "note created", "note", createdNote)

	return &dto.NoteDto{
		ID:              createdNote.ID,
		UserId:          createdNote.UserID,
		Title:           createdNote.Title,
		Content:         createdNote.Content,
		Tags:            createdNote.Tags,
		CreatedDateTime: createdNote.CreatedDateTime,
		UpdatedDateTime: createdNote.UpdatedDateTime,
	}, nil
}

func (n *NoteService) UpdateNoteById(ctx context.Context, note *dto.NoteRequestDto) (*dto.NoteDto, error) {
	updateNote, err := n.noteRepo.GetNoteByID(ctx, note.ID)
	if err != nil {
		return nil, err
	}

	if note.Title != "" {
		updateNote.Title = note.Title
	}

	if note.Content != "" {
		updateNote.Content = note.Content
	}

	var updateTags string
	if len(note.Tags) != 0 {
		updateNote.Tags = note.Tags
		updateTags = strings.Join(updateNote.Tags, ",")
	} else {
		updateTags = strings.Join(updateNote.Tags, ",")
	}

	requestNote, err := domain.ReConstractNote(updateNote.ID, updateNote.UserID, updateNote.Title, updateNote.Content, updateTags, updateNote.CreatedDateTime, updateNote.UpdatedDateTime)
	if err != nil {
		return nil, err
	}

	if err := n.noteRepo.Save(ctx, requestNote); err != nil {
		return nil, err
	}

	updatedNote, err := n.noteRepo.GetNoteByID(ctx, note.ID)
	if err != nil {
		return nil, err
	}

	return &dto.NoteDto{
		ID:              updatedNote.ID,
		UserId:          updatedNote.UserID,
		Title:           updatedNote.Title,
		Content:         updatedNote.Content,
		Tags:            updatedNote.Tags,
		CreatedDateTime: updatedNote.CreatedDateTime,
		UpdatedDateTime: updatedNote.UpdatedDateTime,
	}, nil
}

func (n *NoteService) DeleteNoteById(ctx context.Context, id string) error {
	return n.noteRepo.Delete(ctx, id)
}
