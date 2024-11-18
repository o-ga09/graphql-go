package service

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/o-ga09/graphql-go/internal/domain"
	"github.com/o-ga09/graphql-go/internal/domain/repository"
	mock "github.com/o-ga09/graphql-go/internal/domain/repository/moq"
	"github.com/o-ga09/graphql-go/internal/service/dto"
)

func TestNewNoteService(t *testing.T) {
	type args struct {
		noteRepo repository.NoteRepository
	}
	tests := []struct {
		name string
		args args
		want *NoteService
	}{
		{name: "TestNewNoteService", args: args{noteRepo: nil}, want: &NoteService{noteRepo: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNoteService(tt.args.noteRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNoteService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoteService_FetchNotes(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	res := []*domain.Note{
		{ID: "1", Title: "title1", Content: "content1", Tags: []string{"tag1", "tag2", "tag3"}, CreatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), UpdatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), UserID: "1"},
		{ID: "2", Title: "title2", Content: "content2", Tags: []string{"tag1", "tag2", "tag3"}, CreatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), UpdatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), UserID: "1"},
		{ID: "3", Title: "title3", Content: "content3", Tags: []string{"tag1", "tag2", "tag3"}, CreatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), UpdatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), UserID: "1"},
	}
	expected := []*dto.NoteDto{
		{ID: "1", Title: "title1", Content: "content1", Tags: []string{"tag1", "tag2", "tag3"}, CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00", UserId: "1"},
		{ID: "2", Title: "title2", Content: "content2", Tags: []string{"tag1", "tag2", "tag3"}, CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00", UserId: "1"},
		{ID: "3", Title: "title3", Content: "content3", Tags: []string{"tag1", "tag2", "tag3"}, CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00", UserId: "1"},
	}
	mockedNoteRepository := &mock.NoteRepositoryMock{
		GetNotesFunc: func(contextMoqParam context.Context, s string) ([]*domain.Note, error) {
			return res, nil
		},
	}
	type fields struct {
		noteRepo repository.NoteRepository
	}
	type args struct {
		ctx    context.Context
		userId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*dto.NoteDto
		wantErr bool
	}{
		{name: "TestFetchNotes", fields: fields{noteRepo: mockedNoteRepository}, args: args{ctx: ctx, userId: "1"}, want: expected, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := &NoteService{
				noteRepo: tt.fields.noteRepo,
			}
			got, err := n.FetchNotes(tt.args.ctx, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("NoteService.FetchNotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoteService.FetchNotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoteService_FetchNoteById(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	res := []*domain.Note{
		{ID: "1", Title: "title1", Content: "content1", Tags: []string{"tag1", "tag2", "tag3"}, CreatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), UpdatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), UserID: "1"},
		{ID: "2", Title: "title2", Content: "content2", Tags: []string{"tag1", "tag2", "tag3"}, CreatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), UpdatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), UserID: "1"},
		{ID: "3", Title: "title3", Content: "content3", Tags: []string{"tag1", "tag2", "tag3"}, CreatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), UpdatedDateTime: time.Date(2024, 8, 15, 0, 0, 0, 0, time.UTC), UserID: "1"},
	}
	expected := &dto.NoteDto{ID: "1", Title: "title1", Content: "content1", Tags: []string{"tag1", "tag2", "tag3"}, CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00", UserId: "1"}
	mockedNoteRepository := &mock.NoteRepositoryMock{
		GetNoteByIDFunc: func(contextMoqParam context.Context, id string) (*domain.Note, error) {
			return res[0], nil
		},
	}
	type fields struct {
		noteRepo repository.NoteRepository
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.NoteDto
		wantErr bool
	}{
		{name: "TestFetchNoteById", fields: fields{noteRepo: mockedNoteRepository}, args: args{ctx: ctx, id: "1"}, want: expected, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := &NoteService{
				noteRepo: tt.fields.noteRepo,
			}
			got, err := n.FetchNoteById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("NoteService.FetchNoteById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoteService.FetchNoteById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoteService_CreateNote(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	mockedNoteRepository := &mock.NoteRepositoryMock{
		CreateNoteFunc: func(contextMoqParam context.Context, note *domain.Note) error {
			return nil
		},
		GetNoteByIDFunc: func(contextMoqParam context.Context, id string) (*domain.Note, error) {
			return &domain.Note{ID: "1", Title: "title1", Content: "content1"}, nil
		},
	}
	type fields struct {
		noteRepo repository.NoteRepository
	}
	type args struct {
		ctx  context.Context
		note *domain.Note
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Note
		wantErr bool
	}{
		{name: "TestCreateNote", fields: fields{noteRepo: mockedNoteRepository}, args: args{ctx: ctx, note: &domain.Note{ID: "1", Title: "title1", Content: "content1"}}, want: &domain.Note{ID: "1", Title: "title1", Content: "content1"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := &NoteService{
				noteRepo: tt.fields.noteRepo,
			}
			got, err := n.CreateNote(tt.args.ctx, tt.args.note)
			if (err != nil) != tt.wantErr {
				t.Errorf("NoteService.CreateNote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoteService.CreateNote() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoteService_UpdateNoteById(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	res := []*domain.Note{
		{ID: "1", Title: "title1", Content: "content1"},
		{ID: "2", Title: "title2", Content: "content2"},
		{ID: "3", Title: "title3", Content: "content3"},
	}
	mockedNoteRepository := &mock.NoteRepositoryMock{
		UpdateNoteByIDFunc: func(contextMoqParam context.Context, id string, note *domain.Note) error {
			return nil
		},
		GetNoteByIDFunc: func(contextMoqParam context.Context, id string) (*domain.Note, error) {
			return res[0], nil
		},
	}
	type fields struct {
		noteRepo repository.NoteRepository
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
		want    *domain.Note
		wantErr bool
	}{
		{name: "TestUpdateNoteById", fields: fields{noteRepo: mockedNoteRepository}, args: args{ctx: ctx, id: "1", note: &domain.Note{ID: "1", Title: "title1", Content: "content1"}}, want: &domain.Note{ID: "1", Title: "title1", Content: "content1"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := &NoteService{
				noteRepo: tt.fields.noteRepo,
			}
			got, err := n.UpdateNoteById(tt.args.ctx, tt.args.id, tt.args.note)
			if (err != nil) != tt.wantErr {
				t.Errorf("NoteService.UpdateNoteById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoteService.UpdateNoteById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoteService_DeleteNoteById(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	mockedNoteRepository := &mock.NoteRepositoryMock{
		DeleteNoteByIDFunc: func(contextMoqParam context.Context, id string) error {
			return nil
		},
	}
	type fields struct {
		noteRepo repository.NoteRepository
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
		{name: "TestDeleteNoteById", fields: fields{noteRepo: mockedNoteRepository}, args: args{ctx: ctx, id: "1"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n := &NoteService{
				noteRepo: tt.fields.noteRepo,
			}
			if err := n.DeleteNoteById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("NoteService.DeleteNoteById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
