package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/o-ga09/graphql-go/domain"
	"github.com/o-ga09/graphql-go/graph/model"
	"github.com/o-ga09/graphql-go/pkg/date"
)

// CreateNote is the resolver for the createNote field.
func (r *mutationResolver) CreateNote(ctx context.Context, input model.NewNote) (*model.EditedNote, error) {
	createdNote, err := r.NoteService.CreateNote(ctx, &domain.Note{
		Title:   input.Title,
		Content: input.Content,
		Tags:    input.Tags,
	})
	if err != nil {
		return nil, err
	}
	id, err := uuid.Parse(createdNote.ID)
	if err != nil {
		return nil, err
	}
	return &model.EditedNote{
		ID: id.String(),
	}, nil
}

// UpdateNote is the resolver for the updateNote field.
func (r *mutationResolver) UpdateNote(ctx context.Context, input model.UpdateNote) (*model.EditedNote, error) {
	tags := make([]string, len(input.Tags))
	for i, tag := range input.Tags {
		tags[i] = *tag
	}
	updatedNote, err := r.NoteService.UpdateNoteById(ctx, input.ID, &domain.Note{
		Title:   *input.Title,
		Content: *input.Content,
		Tags:    tags,
	})
	if err != nil {
		return nil, err
	}
	id, err := uuid.Parse(updatedNote.ID)
	if err != nil {
		return nil, err
	}
	return &model.EditedNote{
		ID: id.String(),
	}, nil
}

// DeleteNote is the resolver for the deleteNote field.
func (r *mutationResolver) DeleteNote(ctx context.Context, input model.DeleteNote) (*model.DeletedNote, error) {
	err := r.NoteService.DeleteNoteById(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	return &model.DeletedNote{
		ID: input.ID,
	}, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.EditedUser, error) {
	createdUser, err := r.UserService.CreateUser(ctx, &domain.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Address:   input.Address,
		BirthDay:  input.BirthDay,
		Password:  input.Password,
	})
	if err != nil {
		return nil, err
	}
	return &model.EditedUser{
		ID: createdUser.ID,
	}, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.EditedUser, error) {
	updatedUser, err := r.UserService.UpdateUserById(ctx, input.ID, &domain.User{
		FirstName: *input.FirstName,
		LastName:  *input.LastName,
		Email:     *input.Email,
		Address:   *input.Address,
		BirthDay:  *input.BirthDay,
		Password:  *input.Password,
	})
	if err != nil {
		return nil, err
	}
	return &model.EditedUser{
		ID: updatedUser.ID,
	}, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, input model.DeleteUser) (*model.DeletedUser, error) {
	err := r.UserService.DeleteUserById(ctx, input.ID)
	if err != nil {
		return nil, err
	}
	return &model.DeletedUser{
		ID: input.ID,
	}, nil
}

// Notes is the resolver for the notes field.
func (r *queryResolver) Notes(ctx context.Context, input string) ([]*model.Note, error) {
	note, err := r.NoteService.FetchNotes(ctx, input)
	if err != nil {
		return nil, err
	}
	res := []*model.Note{}
	userid, _ := uuid.NewUUID()
	user := &model.User{
		ID:              userid,
		FirstName:       "",
		LastName:        "",
		Email:           "",
		Address:         "",
		BirthDay:        "",
		Password:        "",
		Sex:             0,
		CreatedDateTime: time.Now(),
		UpdatedDateTime: time.Now(),
	}
	for _, n := range note {
		id, err := uuid.Parse(n.ID)
		if err != nil {
			return nil, err
		}
		createdAt, err := date.TimeToString(n.CreatedDateTime)
		if err != nil {
			return nil, err
		}
		updatedAt, err := date.TimeToString(n.UpdatedDateTime)
		if err != nil {
			return nil, err
		}
		res = append(res, &model.Note{
			ID:              id,
			Title:           n.Title,
			Content:         n.Content,
			Tags:            n.Tags,
			User:            user,
			CreatedDateTime: createdAt,
			UpdatedDateTime: updatedAt,
		})
	}
	return res, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	user, err := r.UserService.FetchUsers(ctx)
	if err != nil {
		return nil, err
	}
	res := []*model.User{}
	for _, u := range user {
		id, err := uuid.Parse(u.ID)
		if err != nil {
			return nil, err
		}
		res = append(res, &model.User{
			ID:              id,
			FirstName:       u.FirstName,
			LastName:        u.LastName,
			Email:           u.Email,
			Address:         u.Address,
			BirthDay:        u.BirthDay,
			Password:        u.Password,
			Sex:             u.Sex,
			CreatedDateTime: u.CreatedDateTime,
			UpdatedDateTime: u.UpdatedDateTime,
		})
	}
	return res, nil
}

// NoteByID is the resolver for the noteById field.
func (r *queryResolver) NoteByID(ctx context.Context, input string) (*model.Note, error) {
	res, err := r.NoteService.FetchNoteById(ctx, input)
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(res.ID)
	if err != nil {
		return nil, err
	}
	createdAt, err := date.TimeToString(res.CreatedDateTime)
	if err != nil {
		return nil, err
	}
	updatedAt, err := date.TimeToString(res.UpdatedDateTime)
	if err != nil {
		return nil, err
	}
	return &model.Note{
		ID:              id,
		Title:           res.Title,
		Content:         res.Content,
		Tags:            res.Tags,
		User:            &model.User{},
		CreatedDateTime: createdAt,
		UpdatedDateTime: updatedAt,
	}, nil
}

// UserByID is the resolver for the userById field.
func (r *queryResolver) UserByID(ctx context.Context, input string) (*model.User, error) {
	return &model.User{
		ID:              uuid.New(),
		FirstName:       "firstName",
		LastName:        "lastName",
		Email:           "email",
		Address:         "address",
		BirthDay:        "birthDay",
		Password:        "password",
		Sex:             0,
		CreatedDateTime: time.Now(),
		UpdatedDateTime: time.Now(),
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
