package dao

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/o-ga09/graphql-go/internal/db/db"
	"github.com/o-ga09/graphql-go/internal/domain"
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
		{ID: "1", UserID: "1", Title: "title1", Content: "content1", Tags: []string{"tags1", "tags2", "tags3"}, CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
	}

	expected := db.GetNotesRow{
		ID:        1,
		NoteID:    "1",
		Title:     "title1",
		Tags:      "tags1,tags2,tags3",
		Content:   "content1",
		CreatedAt: sql.NullTime{Time: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), Valid: true},
		UserID:    "1",
	}

	rows := sqlmock.NewRows([]string{"id", "notes.note_id", "title", "tags", "content", "created_at", "updated_at", "user_id"})
	rows.AddRow(expected.ID, expected.NoteID, expected.Title, expected.Tags, expected.Content, expected.CreatedAt.Time, expected.UpdatedAt.Time, expected.UserID)
	mock.ExpectQuery(`SELECT id, notes.note_id, title, tags, content, created_at, updated_at, user_id FROM notes JOIN user_notes ON notes.note_id = user_notes.note_id WHERE user_notes.user_id = \? AND delete_at IS NULL ORDER BY created_at DESC`).WillReturnRows(rows)

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
		{ID: "1", UserID: "1", Title: "title1", Content: "content1", Tags: []string{"tags1", "tags2", "tags3"}, CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
	}

	expected := db.GetNoteRow{
		ID:        1,
		NoteID:    "1",
		Title:     "title1",
		Tags:      "tags1,tags2,tags3",
		Content:   "content1",
		CreatedAt: sql.NullTime{Time: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), Valid: true},
		UserID:    "1",
	}

	rows := sqlmock.NewRows([]string{"id", "note_id", "title", "tags", "content", "created_at", "updated_at", "user_id"})
	rows.AddRow(expected.ID, expected.NoteID, expected.Title, expected.Tags, expected.Content, expected.CreatedAt.Time, expected.UpdatedAt.Time, expected.UserID)
	mock.ExpectQuery(`SELECT id, notes.note_id, title, tags, content, created_at, updated_at, user_id FROM notes JOIN user_notes ON notes.note_id = user_notes.note_id WHERE user_notes.note_id = \? AND delete_at IS NULL LIMIT 1`).WithArgs("1").WillReturnRows(rows)

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

func Test_noteDao_Save_Create(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	dbmock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	res := []*domain.Note{
		{ID: "1", UserID: "1", Title: "title1", Content: "content1", Tags: []string{"tags1", "tags2", "tags3"}, CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
	}

	mock.ExpectExec("INSERT INTO notes").WithArgs("1", "title1", "tags1,tags2,tags3", "content1").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO user_notes").WithArgs("1", "1").WillReturnResult(sqlmock.NewResult(1, 1))

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
		{name: "TestCreateNote", fields: fields{query: db.New(dbmock)}, args: args{ctx: ctx, note: res[0]}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := &noteDao{
				query: tt.fields.query,
			}
			if err := n.Save(tt.args.ctx, tt.args.note); (err != nil) != tt.wantErr {
				t.Errorf("noteDao.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_noteDao_Save_Update(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	dbmock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	res := []*domain.Note{
		{ID: "1", Title: "title1", Content: "content1", Tags: []string{"tags1", "tags2", "tags3"}, CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
	}

	expected := db.GetNoteRow{
		ID:        1,
		NoteID:    "1",
		Title:     "title1",
		Tags:      "tags1,tags2,tags3",
		Content:   "content1",
		CreatedAt: sql.NullTime{Time: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), Valid: true},
		UserID:    "1",
	}

	rows := sqlmock.NewRows([]string{"id", "notes.note_id", "title", "tags", "content", "created_at", "updated_at", "user_id"})
	rows.AddRow(expected.ID, expected.NoteID, expected.Title, expected.Tags, expected.Content, expected.CreatedAt.Time, expected.UpdatedAt.Time, expected.UserID)
	mock.ExpectQuery(`SELECT id, notes.note_id, title, tags, content, created_at, updated_at, user_id FROM notes JOIN user_notes ON notes.note_id = user_notes.note_id WHERE user_notes.note_id = \? AND delete_at IS NULL LIMIT 1`).WithArgs("1").WillReturnRows(rows)

	arg := db.UpdateNoteParams{
		NoteID:  "1",
		Title:   "title1",
		Tags:    "tags1,tags2,tags3",
		Content: "content1",
	}

	mock.ExpectExec("UPDATE notes").WithArgs(arg.Title, arg.Tags, arg.Content, arg.NoteID).WillReturnResult(sqlmock.NewResult(1, 1))

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
		{name: "TestUpdateNote", fields: fields{query: db.New(dbmock)}, args: args{ctx: ctx, note: res[0]}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := &noteDao{
				query: tt.fields.query,
			}
			if err := n.Save(tt.args.ctx, tt.args.note); (err != nil) != tt.wantErr {
				t.Errorf("noteDao.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_noteDao_Delete(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	dbmock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	mock.ExpectExec("UPDATE notes").WithArgs("1").WillReturnResult(sqlmock.NewResult(1, 1))

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
			if err := n.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("noteDao.DeleteNoteByID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
