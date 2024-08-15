package dao

import (
	"context"
	"database/sql"
	"strings"

	"github.com/o-ga09/graphql-go/db/db"
	"github.com/o-ga09/graphql-go/domain"
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

func (n *noteDao) GetNotes(ctx context.Context) ([]*domain.Note, error) {
	notes, err := n.query.GetNotes(ctx)
	if err != nil {
		return nil, err
	}
	res := []*domain.Note{}
	for _, note := range notes {
		r, err := domain.NewNote(note.NoteID, note.Title, note.Content, note.Tags, note.CreatedAt.Time.Format("2006-01-02 15:04:05"), note.UpdatedAt.Time.Format("2006-01-02 15:04:05"))
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}

func (n *noteDao) GetNoteByID(ctx context.Context, id string) (*domain.Note, error) {
	note, err := n.query.GetNote(ctx, id)
	if err != nil {
		return nil, err
	}
	return domain.NewNote(note.NoteID, note.Title, note.Content, note.Tags, note.CreatedAt.Time.Format("2006-01-02 15:04:05"), note.UpdatedAt.Time.Format("2006-01-02 15:04:05"))
}

func (n *noteDao) CreateNote(ctx context.Context, note *domain.Note) error {
	tags := strings.Join(note.Tags, ",")
	_, err := n.query.CreateNote(ctx, db.CreateNoteParams{
		NoteID:  note.ID,
		Title:   note.Title,
		Tags:    tags,
		Content: note.Content,
	})
	return err
}

func (n *noteDao) UpdateNoteByID(ctx context.Context, id string, note *domain.Note) error {
	tags := strings.Join(note.Tags, ",")
	err := n.query.UpdateNote(ctx, db.UpdateNoteParams{
		NoteID:  id,
		Title:   note.Title,
		Tags:    tags,
		Content: note.Content,
	})
	return err
}

func (n *noteDao) DeleteNoteByID(ctx context.Context, id string) error {
	err := n.query.DeleteNote(ctx, id)
	return err
}
