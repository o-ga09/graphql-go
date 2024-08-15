package repository

import "github.com/o-ga09/graphql-go/domain"

type UserRepository interface {
	GetUsers() ([]*domain.User, error)
	GetUserById(id string) (*domain.User, error)
	CreateUser(user *domain.User) error
	UpdateUserById(id string, user *domain.User) error
	DeleteUserById(id string) error
}
