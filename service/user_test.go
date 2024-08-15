package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/o-ga09/graphql-go/domain"
	"github.com/o-ga09/graphql-go/domain/repository"
	mock "github.com/o-ga09/graphql-go/domain/repository/moq"
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
		{ID: "1", FirstName: "firstName1", LastName: "lastName1", Email: "email1", Address: "address1", BirthDay: "birthDay1", Password: "password", Sex: 0},
		{ID: "2", FirstName: "firstName2", LastName: "lastName2", Email: "email2", Address: "address2", BirthDay: "birthDay2", Password: "password", Sex: 1},
		{ID: "3", FirstName: "firstName3", LastName: "lastName3", Email: "email3", Address: "address3", BirthDay: "birthDay3", Password: "password", Sex: 0},
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
		want    []*domain.User
		wantErr bool
	}{
		{name: "TestFetchUsers", fields: fields{userRepo: mockedUserRepository}, args: args{ctx: ctx}, want: res, wantErr: false},
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
		{ID: "1", FirstName: "firstName1", LastName: "lastName1", Email: "email1", Address: "address1", BirthDay: "birthDay1", Password: "password", Sex: 0},
		{ID: "2", FirstName: "firstName2", LastName: "lastName2", Email: "email2", Address: "address2", BirthDay: "birthDay2", Password: "password", Sex: 1},
		{ID: "3", FirstName: "firstName3", LastName: "lastName3", Email: "email3", Address: "address3", BirthDay: "birthDay3", Password: "password", Sex: 0},
	}
	mockedUserRepository := &mock.UserRepositoryMock{
		GetUserByIdFunc: func(contextMoqParam context.Context, id string) (*domain.User, error) {
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
		want    *domain.User
		wantErr bool
	}{
		{name: "TestFetchUserById", fields: fields{userRepo: mockedUserRepository}, args: args{ctx: ctx, id: "1"}, want: res[0], wantErr: false},
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
	res := []*domain.User{
		{ID: "1", FirstName: "firstName1", LastName: "lastName1", Email: "email1", Address: "address1", BirthDay: "birthDay1", Password: "password", Sex: 0},
		{ID: "2", FirstName: "firstName2", LastName: "lastName2", Email: "email2", Address: "address2", BirthDay: "birthDay2", Password: "password", Sex: 1},
		{ID: "3", FirstName: "firstName3", LastName: "lastName3", Email: "email3", Address: "address3", BirthDay: "birthDay3", Password: "password", Sex: 0},
	}
	mockedUserRepository := &mock.UserRepositoryMock{
		CreateUserFunc: func(contextMoqParam context.Context, user *domain.User) error {
			return nil
		},
		GetUserByIdFunc: func(contextMoqParam context.Context, id string) (*domain.User, error) {
			return res[0], nil
		},
	}
	type fields struct {
		userRepo repository.UserRepository
	}
	type args struct {
		ctx  context.Context
		user *domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.User
		wantErr bool
	}{
		{name: "TestCreateUser", fields: fields{userRepo: mockedUserRepository}, args: args{ctx: ctx, user: res[0]}, want: res[0], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := &UserService{
				userRepo: tt.fields.userRepo,
			}
			got, err := u.CreateUser(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UpdateUserById(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	res := []*domain.User{
		{ID: "1", FirstName: "firstName1", LastName: "lastName1", Email: "email1", Address: "address1", BirthDay: "birthDay1", Password: "password", Sex: 0},
		{ID: "2", FirstName: "firstName2", LastName: "lastName2", Email: "email2", Address: "address2", BirthDay: "birthDay2", Password: "password", Sex: 1},
		{ID: "3", FirstName: "firstName3", LastName: "lastName3", Email: "email3", Address: "address3", BirthDay: "birthDay3", Password: "password", Sex: 0},
	}
	mockedUserRepository := &mock.UserRepositoryMock{
		UpdateUserByIdFunc: func(contextMoqParam context.Context, id string, user *domain.User) error {
			return nil
		},
		GetUserByIdFunc: func(contextMoqParam context.Context, id string) (*domain.User, error) {
			return res[0], nil
		},
	}
	type fields struct {
		userRepo repository.UserRepository
	}
	type args struct {
		ctx  context.Context
		id   string
		user *domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.User
		wantErr bool
	}{
		{name: "TestUpdateUserById", fields: fields{userRepo: mockedUserRepository}, args: args{ctx: ctx, id: "1", user: res[0]}, want: res[0], wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u := &UserService{
				userRepo: tt.fields.userRepo,
			}
			got, err := u.UpdateUserById(tt.args.ctx, tt.args.id, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.UpdateUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.UpdateUserById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_DeleteUserById(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	mockedUserRepository := &mock.UserRepositoryMock{
		DeleteUserByIdFunc: func(contextMoqParam context.Context, id string) error {
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
