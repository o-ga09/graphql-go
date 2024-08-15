package repository

import (
	"context"

	"github.com/o-ga09/graphql-go/domain"
)

//go:generate moq -out moq/user_repository_mock.go -pkg mock . UserRepository
type UserRepository interface {
	GetUsers(context.Context) ([]*domain.User, error)
	GetUserById(ctx context.Context, id string) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
	UpdateUserById(ctx context.Context, id string, user *domain.User) error
	DeleteUserById(ctx context.Context, id string) error
}
