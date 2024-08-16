package dao

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/o-ga09/graphql-go/db/db"
	"github.com/o-ga09/graphql-go/domain"
)

func TestNewNoteDao(t *testing.T) {
	t.Parallel()
	dbmock, _, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		d *sql.DB
	}
	tests := []struct {
		name string
		args args
		want *noteDao
	}{
		{name: "success", args: args{d: dbmock}, want: &noteDao{query: db.New(dbmock)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NewNoteDao(tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNoteDao() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_noteDao_GetNotes(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	dbmock, mock, err := sqlmock.New()
	res := []*domain.Note{
		{ID: "1", Title: "title1", Content: "content1", Tags: []string{"tags1", "tags2", "tags3"}, CreatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), UpdatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC)},
	}

	expected := db.Note{
		ID:        1,
		NoteID:    "1",
		Title:     "title1",
		Tags:      "tags1,tags2,tags3",
		Content:   "content1",
		CreatedAt: sql.NullTime{Time: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), Valid: true},
	}

	rows := sqlmock.NewRows([]string{"id", "note_id", "title", "tags", "content", "created_at", "updated_at"})
	rows.AddRow(expected.ID, expected.NoteID, expected.Title, expected.Tags, expected.Content, expected.CreatedAt.Time, expected.UpdatedAt.Time)
	mock.ExpectQuery("SELECT id, note_id, title, tags, content, created_at, updated_at FROM notes ORDER BY created_at DESC").WillReturnRows(rows)

	if err != nil {
		t.Fatal(err)
	}
	type fields struct {
		query *db.Queries
	}
	type args struct {
		ctx    context.Context
		userId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*domain.Note
		wantErr bool
	}{
		{name: "success", fields: fields{query: db.New(dbmock)}, args: args{ctx: ctx, userId: "1"}, want: res},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := &noteDao{
				query: tt.fields.query,
			}
			got, err := n.GetNotes(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("noteDao.GetNotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("noteDao.GetNotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_noteDao_GetNoteByID(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	dbmock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	res := []*domain.Note{
		{ID: "1", Title: "title1", Content: "content1", Tags: []string{"tags1", "tags2", "tags3"}, CreatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), UpdatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC)},
	}

	expected := db.Note{
		ID:        1,
		NoteID:    "1",
		Title:     "title1",
		Tags:      "tags1,tags2,tags3",
		Content:   "content1",
		CreatedAt: sql.NullTime{Time: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), Valid: true},
	}

	rows := sqlmock.NewRows([]string{"id", "note_id", "title", "tags", "content", "created_at", "updated_at"})
	rows.AddRow(expected.ID, expected.NoteID, expected.Title, expected.Tags, expected.Content, expected.CreatedAt.Time, expected.UpdatedAt.Time)
	mock.ExpectQuery("SELECT id, note_id, title, tags, content, created_at, updated_at FROM notes WHERE note_id = ?").WithArgs("1").WillReturnRows(rows)

	type fields struct {
		query *db.Queries
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Note
		wantErr bool
	}{
		{name: "success", fields: fields{query: db.New(dbmock)}, args: args{ctx: ctx, id: "1"}, want: res[0]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := &noteDao{
				query: tt.fields.query,
			}
			got, err := n.GetNoteByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("noteDao.GetNoteByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("noteDao.GetNoteByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_noteDao_CreateNote(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	dbmock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	res := []*domain.Note{
		{ID: "1", Title: "title1", Content: "content1", Tags: []string{"tags1", "tags2", "tags3"}, CreatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), UpdatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC)},
	}

	mock.ExpectExec("INSERT INTO notes").WithArgs("1", "title1", "tags1,tags2,tags3", "content1").WillReturnResult(sqlmock.NewResult(1, 1))

	type fields struct {
		query *db.Queries
	}
	type args struct {
		ctx  context.Context
		note *domain.Note
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "success", fields: fields{query: db.New(dbmock)}, args: args{ctx: ctx, note: res[0]}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := &noteDao{
				query: tt.fields.query,
			}
			if err := n.CreateNote(tt.args.ctx, tt.args.note); (err != nil) != tt.wantErr {
				t.Errorf("noteDao.CreateNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_noteDao_UpdateNoteByID(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	dbmock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	res := []*domain.Note{
		{ID: "1", Title: "title1", Content: "content1", Tags: []string{"tags1", "tags2", "tags3"}, CreatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), UpdatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC)},
	}

	mock.ExpectExec("UPDATE notes").WithArgs("title1", "tags1,tags2,tags3", "content1", "1").WillReturnResult(sqlmock.NewResult(1, 1))

	type fields struct {
		query *db.Queries
	}
	type args struct {
		ctx  context.Context
		id   string
		note *domain.Note
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "success", fields: fields{query: db.New(dbmock)}, args: args{ctx: ctx, id: "1", note: res[0]}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := &noteDao{
				query: tt.fields.query,
			}
			if err := n.UpdateNoteByID(tt.args.ctx, tt.args.id, tt.args.note); (err != nil) != tt.wantErr {
				t.Errorf("noteDao.UpdateNoteByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_noteDao_DeleteNoteByID(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	dbmock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectExec("DELETE FROM notes").WithArgs("1").WillReturnResult(sqlmock.NewResult(1, 1))

	type fields struct {
		query *db.Queries
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "success", fields: fields{query: db.New(dbmock)}, args: args{ctx: ctx, id: "1"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := &noteDao{
				query: tt.fields.query,
			}
			if err := n.DeleteNoteByID(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("noteDao.DeleteNoteByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
