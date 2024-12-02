package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/o-ga09/graphql-go/internal/domain"
	"github.com/o-ga09/graphql-go/internal/domain/repository"
	mock "github.com/o-ga09/graphql-go/internal/domain/repository/moq"
	"github.com/o-ga09/graphql-go/internal/service/dto"
)

func TestNewUserService(t *testing.T) {
	t.Parallel()
	type args struct {
		userRepo repository.UserRepository
	}
	tests := []struct {
		name string
		args args
		want *UserService
	}{
		{name: "TestNewUserService", args: args{userRepo: nil}, want: &UserService{userRepo: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := NewUserService(tt.args.userRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_FetchUsers(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	res := []*domain.User{
		{ID: "1", UserName: "userName1", DisplayName: "displayName1", CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
		{ID: "2", UserName: "userName2", DisplayName: "displayName2", CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
		{ID: "3", UserName: "userName3", DisplayName: "displayName3", CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
	}

	expected := []*dto.UserDto{
		{ID: "1", UserName: "userName1", DisplayName: "displayName1", CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
		{ID: "2", UserName: "userName2", DisplayName: "displayName2", CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
		{ID: "3", UserName: "userName3", DisplayName: "displayName3", CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
	}

	mockedUserRepository := &mock.UserRepositoryMock{
		GetUsersFunc: func(contextMoqParam context.Context) ([]*domain.User, error) {
			return res, nil
		},
	}
	type fields struct {
		userRepo repository.UserRepository
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*dto.UserDto
		wantErr bool
	}{
		{name: "TestFetchUsers", fields: fields{userRepo: mockedUserRepository}, args: args{ctx: ctx}, want: expected, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := &UserService{
				userRepo: tt.fields.userRepo,
			}
			got, err := u.FetchUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.FetchUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.FetchUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_FetchUserById(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	res := []*domain.User{
		{ID: "1", UserName: "userName1", DisplayName: "displayName1", CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
		{ID: "2", UserName: "userName2", DisplayName: "displayName2", CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
		{ID: "3", UserName: "userName3", DisplayName: "displayName3", CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
	}
	expected := &dto.UserDto{
		ID:              "1",
		UserName:        "userName1",
		DisplayName:     "displayName1",
		CreatedDateTime: "2024-08-15 00:00:00",
		UpdatedDateTime: "2024-08-15 00:00:00",
	}
	mockedUserRepository := &mock.UserRepositoryMock{
		GetUserByIDFunc: func(ctx context.Context, id string) (*domain.User, error) {
			return res[0], nil
		},
	}
	type fields struct {
		userRepo repository.UserRepository
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.UserDto
		wantErr bool
	}{
		{name: "TestFetchUserById", fields: fields{userRepo: mockedUserRepository}, args: args{ctx: ctx, id: "1"}, want: expected, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := &UserService{
				userRepo: tt.fields.userRepo,
			}
			got, err := u.FetchUserById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.FetchUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.FetchUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_CreateUser(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	mockres := []*domain.User{
		{ID: "1", UserName: "userName1", DisplayName: "displayName1", CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
		{ID: "2", UserName: "userName2", DisplayName: "displayName2", CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
		{ID: "3", UserName: "userName3", DisplayName: "displayName3", CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
	}
	argUser := []*dto.UserReqsutDto{
		{UserId: "1", UserName: "userName1", DisplayName: "displayName1"},
		{UserId: "2", UserName: "userName2", DisplayName: "displayName2"},
		{UserId: "3", UserName: "userName3", DisplayName: "displayName3"},
	}
	expected := []*dto.UserDto{
		{ID: "1", UserName: "userName1", DisplayName: "displayName1", CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
		{ID: "2", UserName: "userName2", DisplayName: "displayName2", CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
		{ID: "3", UserName: "userName3", DisplayName: "displayName3", CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
	}
	mockedUserRepository := &mock.UserRepositoryMock{
		SaveFunc: func(ctx context.Context, user *domain.User) error {
			return nil
		},
		GetUserByIDFunc: func(ctx context.Context, id string) (*domain.User, error) {
			return mockres[0], nil
		},
	}
	type fields struct {
		userRepo repository.UserRepository
	}
	type args struct {
		ctx  context.Context
		user *dto.UserReqsutDto
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.UserDto
		wantErr bool
	}{
		{name: "TestCreateUser", fields: fields{userRepo: mockedUserRepository}, args: args{ctx: ctx, user: argUser[0]}, want: expected[0], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := &UserService{
				userRepo: tt.fields.userRepo,
			}
			err := u.CreateUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserService_UpdateUserById(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	mockres := []*domain.User{
		{ID: "1", UserName: "userName1", DisplayName: "displayName1", CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
		{ID: "2", UserName: "userName2", DisplayName: "displayName2", CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
		{ID: "3", UserName: "userName3", DisplayName: "displayName3", CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
	}
	argUser := []*dto.UserReqsutDto{
		{UserId: "1", UserName: "userName1", DisplayName: "displayName1"},
		{UserId: "2", UserName: "userName2", DisplayName: "displayName2"},
		{UserId: "3", UserName: "userName3", DisplayName: "displayName3"},
	}
	expected := []*dto.UserDto{
		{ID: "1", UserName: "userName1", DisplayName: "displayName1", CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
		{ID: "2", UserName: "userName2", DisplayName: "displayName2", CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
		{ID: "3", UserName: "userName3", DisplayName: "displayName3", CreatedDateTime: "2024-08-15 00:00:00", UpdatedDateTime: "2024-08-15 00:00:00"},
	}
	mockedUserRepository := &mock.UserRepositoryMock{
		SaveFunc: func(ctx context.Context, user *domain.User) error {
			return nil
		},
		GetUserByIDFunc: func(ctx context.Context, id string) (*domain.User, error) {
			return mockres[0], nil
		},
	}
	type fields struct {
		userRepo repository.UserRepository
	}
	type args struct {
		ctx  context.Context
		id   string
		user *dto.UserReqsutDto
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dto.UserDto
		wantErr bool
	}{
		{name: "TestUpdateUserById", fields: fields{userRepo: mockedUserRepository}, args: args{ctx: ctx, id: "1", user: argUser[0]}, want: expected[0], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := &UserService{
				userRepo: tt.fields.userRepo,
			}
			err := u.UpdateUserById(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.UpdateUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserService_DeleteUserById(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	mockedUserRepository := &mock.UserRepositoryMock{
		DeleteFunc: func(ctx context.Context, id string) error {
			return nil
		},
	}
	type fields struct {
		userRepo repository.UserRepository
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
		{name: "TestDeleteUserById", fields: fields{userRepo: mockedUserRepository}, args: args{ctx: ctx, id: "1"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := &UserService{
				userRepo: tt.fields.userRepo,
			}
			if err := u.DeleteUserById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UserService.DeleteUserById() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
