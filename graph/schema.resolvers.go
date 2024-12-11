package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/o-ga09/graphql-go/graph/model"
	"github.com/o-ga09/graphql-go/internal/domain"
	"github.com/o-ga09/graphql-go/internal/service/dto"
	"github.com/o-ga09/graphql-go/pkg/date"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, userID string, username string, displayname string) (*model.User, error) {
	new := &dto.UserReqsutDto{
		UserId:      userID,
		UserName:    username,
		DisplayName: displayname,
	}
	if err := r.UserService.CreateUser(ctx, new); err != nil {
		return nil, err
	}
	createdUser, err := r.UserService.FetchUserById(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &model.User{
		UserID:      createdUser.ID,
		Username:    createdUser.UserName,
		Displayname: createdUser.DisplayName,
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, userID string, username string, displayname *string) (*model.User, error) {
	updateUser := &dto.UserReqsutDto{
		UserId:      userID,
		UserName:    username,
		DisplayName: *displayname,
	}
	if err := r.UserService.UpdateUserById(ctx, updateUser); err != nil {
		return nil, err
	}

	updatedUser, err := r.UserService.FetchUserById(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &model.User{
		UserID:      updatedUser.ID,
		Username:    updatedUser.UserName,
		Displayname: updatedUser.DisplayName,
	}, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, userID string) (*string, error) {
	if err := r.UserService.DeleteUserById(ctx, userID); err != nil {
		return nil, err
	}
	msg := "Deleted"
	return &msg, nil
}

// CreateNote is the resolver for the createNote field.
func (r *mutationResolver) CreateNote(ctx context.Context, userID string, title string, content string, tags []string) (*model.Note, error) {
	newTags := strings.Join(tags, ",")
	note, err := domain.NewNote(userID, title, content, newTags, time.Now().String(), time.Now().String())
	if err != nil {
		return nil, fmt.Errorf("failed to create note: %w", err)
	}

	createdNote, err := r.NoteService.CreateNote(ctx, note)
	if err != nil {
		return nil, err
	}

	createdTags := []*model.PostTag{}
	for _, t := range createdNote.Tags {
		createdTags = append(createdTags, &model.PostTag{
			Name: t,
		})
	}

	createdDateTime, err := date.TimeToString(createdNote.CreatedDateTime)
	if err != nil {
		return nil, err
	}

	updatedDateTime, err := date.TimeToString(createdNote.UpdatedDateTime)
	if err != nil {
		return nil, err
	}

	return &model.Note{
		NoteID:    createdNote.ID,
		Title:     createdNote.Title,
		Content:   createdNote.Content,
		Tags:      createdTags,
		CreatedAt: createdDateTime,
		UpdatedAt: updatedDateTime,
	}, nil
}

// UpdateNote is the resolver for the updateNote field.
func (r *mutationResolver) UpdateNote(ctx context.Context, noteID string, title *string, content *string, tags []string) (*model.Note, error) {
	note := &dto.NoteRequestDto{
		ID:      noteID,
		Title:   *title,
		Content: *content,
		Tags:    tags,
	}

	updatedNote, err := r.NoteService.UpdateNoteById(ctx, note)
	if err != nil {
		return nil, err
	}

	updatedTags := []*model.PostTag{}
	for _, t := range updatedNote.Tags {
		updatedTags = append(updatedTags, &model.PostTag{
			Name: t,
		})
	}

	createdDateTime, err := date.TimeToString(updatedNote.CreatedDateTime)
	if err != nil {
		return nil, err
	}

	updatedDateTime, err := date.TimeToString(updatedNote.UpdatedDateTime)
	if err != nil {
		return nil, err
	}

	return &model.Note{
		NoteID:    updatedNote.ID,
		Title:     updatedNote.Title,
		Content:   updatedNote.Content,
		Tags:      updatedTags,
		CreatedAt: createdDateTime,
		UpdatedAt: updatedDateTime,
	}, nil
}

// DeleteNote is the resolver for the deleteNote field.
func (r *mutationResolver) DeleteNote(ctx context.Context, noteID string) (*string, error) {
	err := r.NoteService.DeleteNoteById(ctx, noteID)
	if err != nil {
		return nil, err
	}

	msg := "Deleted"
	return &msg, nil
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	user, err := r.UserService.FetchUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.User{
		UserID:      user.ID,
		Username:    user.UserName,
		Displayname: user.DisplayName,
	}, nil
}

// GetUsers is the resolver for the getUsers field.
func (r *queryResolver) GetUsers(ctx context.Context) ([]*model.User, error) {
	users, err := r.UserService.FetchUsers(ctx)
	if err != nil {
		return nil, err
	}

	res := []*model.User{}
	for _, u := range users {
		res = append(res, &model.User{
			UserID:      u.ID,
			Username:    u.UserName,
			Displayname: u.DisplayName,
		})
	}

	return res, nil
}

// GetNotesByUserID is the resolver for the getNotesByUserId field.
func (r *queryResolver) GetNotesByUserID(ctx context.Context, userID string) (*model.Notes, error) {
	note, err := r.NoteService.FetchNotesByUserId(ctx, userID)
	if err != nil {
		return nil, err
	}

	res := []*model.Note{}
	for _, n := range note {
		createdDateTime, err := date.TimeToString(n.CreatedDateTime)
		if err != nil {
			return nil, err
		}
		updatedDateTime, err := date.TimeToString(n.UpdatedDateTime)
		if err != nil {
			return nil, err
		}

		tags := []*model.PostTag{}
		for _, t := range n.Tags {
			tags = append(tags, &model.PostTag{
				Name: t,
			})
		}

		res = append(res, &model.Note{
			NoteID:    n.ID,
			Title:     n.Title,
			Content:   n.Content,
			Tags:      tags,
			CreatedAt: createdDateTime,
			UpdatedAt: updatedDateTime,
		})
	}

	user, err := r.UserService.FetchUserById(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &model.Notes{
		Count: len(res),
		Notes: res,
		Author: &model.User{
			UserID:      user.ID,
			Username:    user.UserName,
			Displayname: user.DisplayName,
		},
	}, nil
}

// GetNoteAll is the resolver for the getNoteAll field.
func (r *queryResolver) GetNoteAll(ctx context.Context) ([]*model.Note, error) {
	notes, err := r.NoteService.FetchNoteAll(ctx)
	if err != nil {
		return nil, err
	}

	res := []*model.Note{}
	for _, n := range notes {
		createdDateTime, err := date.TimeToString(n.CreatedDateTime)
		if err != nil {
			return nil, err
		}
		updatedDateTime, err := date.TimeToString(n.UpdatedDateTime)
		if err != nil {
			return nil, err
		}

		tags := []*model.PostTag{}
		for _, t := range n.Tags {
			tags = append(tags, &model.PostTag{
				Name: t,
			})
		}

		res = append(res, &model.Note{
			NoteID:    n.ID,
			Title:     n.Title,
			Content:   n.Content,
			Tags:      tags,
			CreatedAt: createdDateTime,
			UpdatedAt: updatedDateTime,
		})
	}

	return res, nil
}

// GetNoteByID is the resolver for the getNoteById field.
func (r *queryResolver) GetNoteByID(ctx context.Context, id string) (*model.NoteByAuthor, error) {
	note, err := r.NoteService.FetchNoteById(ctx, id)
	if err != nil {
		return nil, err
	}

	createdDateTime, err := date.TimeToString(note.CreatedDateTime)
	if err != nil {
		return nil, err
	}

	updatedDateTime, err := date.TimeToString(note.UpdatedDateTime)
	if err != nil {
		return nil, err
	}

	tags := []*model.PostTag{}
	for _, t := range note.Tags {
		tags = append(tags, &model.PostTag{
			Name: t,
		})
	}

	user, err := r.UserService.FetchUserById(ctx, note.UserId)
	if err != nil {
		return nil, err
	}

	return &model.NoteByAuthor{
		Note: &model.Note{
			NoteID:    note.ID,
			Title:     note.Title,
			Content:   note.Content,
			CreatedAt: createdDateTime,
			UpdatedAt: updatedDateTime,
			Tags:      tags,
		},
		Author: &model.User{
			UserID:      user.ID,
			Username:    user.UserName,
			Displayname: user.DisplayName,
		},
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
