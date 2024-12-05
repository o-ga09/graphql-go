package dao

import (
	"context"
	"database/sql"
	"log"
	"strings"

	"github.com/o-ga09/graphql-go/internal/db/db"
	"github.com/o-ga09/graphql-go/internal/domain"
)

type noteDao struct {
	query *db.Queries
}

func NewNoteDao(d *sql.DB) *noteDao {
	q := db.New(d)
	return &noteDao{
		query: q,
	}
}

func (n *noteDao) GetNotes(ctx context.Context, userId string) ([]*domain.Note, error) {
	notes, err := n.query.GetNotes(ctx, userId)
	if err != nil {
		return nil, err
	}
	res := []*domain.Note{}
	for _, note := range notes {
		createdDateTime := strings.Replace(note.CreatedAt.String, " +0000 UTC", "", 1)
		updatedDateTime := strings.Replace(note.UpdatedAt.String, " +0000 UTC", "", 1)
		log.Println(createdDateTime)
		r, err := domain.ReConstractNote(note.NoteID, note.UserID, note.Title, note.Content, note.Tags, createdDateTime, updatedDateTime)
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	log.Println(res[0].CreatedDateTime)
	return res, nil
}

func (n *noteDao) GetNoteByID(ctx context.Context, id string) (*domain.Note, error) {
	note, _ := n.query.GetNote(ctx, id)
	createdDateTime := strings.Replace(note.CreatedAt.String, " +0000 UTC", "", 1)
	updatedDateTime := strings.Replace(note.UpdatedAt.String, " +0000 UTC", "", 1)
	return domain.ReConstractNote(note.NoteID, note.UserID, note.Title, note.Content, note.Tags, createdDateTime, updatedDateTime)
}

func (n *noteDao) Save(ctx context.Context, note *domain.Note) error {
	record, _ := n.query.GetNote(ctx, note.ID)
	if record.NoteID != note.ID {
		_, err := n.query.CreateNote(ctx, db.CreateNoteParams{
			NoteID:  note.ID,
			Title:   note.Title,
			Content: note.Content,
			Tags:    strings.Join(note.Tags, ","),
		})
		if err != nil {
			return err
		}

		if err := n.query.CreateUserNote(ctx, db.CreateUserNoteParams{
			UserID: note.UserID,
			NoteID: note.ID,
		}); err != nil {
			return err
		}
	} else {
		err := n.query.UpdateNote(ctx, db.UpdateNoteParams{
			NoteID:  note.ID,
			Title:   note.Title,
			Tags:    strings.Join(note.Tags, ","),
			Content: note.Content,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *noteDao) Delete(ctx context.Context, id string) error {
	err := n.query.DeleteNote(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
