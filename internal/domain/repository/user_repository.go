package repository

import (
	"context"

	"github.com/o-ga09/graphql-go/internal/domain"
)

//go:generate moq -out moq/user_repository_mock.go -pkg mock . UserRepository
type UserRepository interface {
	GetUserByID(ctx context.Context, id string) (*domain.User, error)
	GetUsers(ctx context.Context) ([]*domain.User, error)
	Save(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id string) error
}
