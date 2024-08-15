package service

import (
	"context"

	"github.com/o-ga09/graphql-go/domain"
	"github.com/o-ga09/graphql-go/domain/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) FetchUsers(ctx context.Context) ([]*domain.User, error) {
	users, err := u.userRepo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserService) FetchUserById(ctx context.Context, id string) (*domain.User, error) {
	user, err := u.userRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) CreateUser(ctx context.Context, user *domain.User) error {
	err := u.userRepo.CreateUser(ctx, user)
	return err
}

func (u *UserService) UpdateUserById(ctx context.Context, id string, user *domain.User) error {
	err := u.userRepo.UpdateUserById(ctx, id, user)
	return err
}

func (u *UserService) DeleteUserById(ctx context.Context, id string) error {
	err := u.userRepo.DeleteUserById(ctx, id)
	return err
}
