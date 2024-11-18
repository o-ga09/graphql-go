package service

import (
	"context"

	"github.com/o-ga09/graphql-go/internal/domain"
	"github.com/o-ga09/graphql-go/internal/domain/repository"
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

func (u *UserService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &domain.User{
		ID:              user.ID,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Email:           user.Email,
		Address:         user.Address,
		BirthDay:        user.BirthDay,
		Password:        user.Password,
		Sex:             user.Sex,
		CreatedDateTime: user.CreatedDateTime,
		UpdatedDateTime: user.UpdatedDateTime,
	}, err
}

func (u *UserService) UpdateUserById(ctx context.Context, id string, user *domain.User) (*domain.User, error) {
	err := u.userRepo.UpdateUserById(ctx, id, user)
	if err != nil {
		return nil, err
	}
	updatedUser, err := u.userRepo.GetUserById(ctx, id)
	return updatedUser, err
}

func (u *UserService) DeleteUserById(ctx context.Context, id string) error {
	err := u.userRepo.DeleteUserById(ctx, id)
	return err
}
