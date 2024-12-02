package service

import (
	"context"
	"strings"
	"time"

	"github.com/o-ga09/graphql-go/internal/domain"
	"github.com/o-ga09/graphql-go/internal/domain/repository"
	"github.com/o-ga09/graphql-go/internal/service/dto"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) FetchUsers(ctx context.Context) ([]*dto.UserDto, error) {
	users, err := u.userRepo.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	res := []*dto.UserDto{}
	for _, user := range users {
		r := &dto.UserDto{
			ID:              user.ID,
			UserName:        user.UserName,
			DisplayName:     user.DisplayName,
			CreatedDateTime: user.CreatedDateTime,
			UpdatedDateTime: user.UpdatedDateTime,
		}
		res = append(res, r)
	}
	return res, nil
}

func (u *UserService) FetchUserById(ctx context.Context, id string) (*dto.UserDto, error) {
	user, err := u.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.UserDto{
		ID:              user.ID,
		UserName:        user.UserName,
		DisplayName:     user.DisplayName,
		CreatedDateTime: user.CreatedDateTime,
		UpdatedDateTime: user.UpdatedDateTime,
	}, nil
}

func (u *UserService) CreateUser(ctx context.Context, user *dto.UserReqsutDto) error {
	createdDateTime := strings.Replace(time.Now().String(), " +0000 UTC", "", 1)
	updatedDateTime := strings.Replace(time.Now().String(), " +0000 UTC", "", 1)
	createdUser, err := domain.NewUser(user.UserId, user.UserName, user.DisplayName, createdDateTime, updatedDateTime)
	if err != nil {
		return err
	}
	if err := u.userRepo.Save(ctx, createdUser); err != nil {
		return err
	}
	return nil
}

func (u *UserService) UpdateUserById(ctx context.Context, user *dto.UserReqsutDto) error {
	createdDateTime := strings.Replace(time.Now().String(), " +0000 UTC", "", 1)
	updatedDateTime := strings.Replace(time.Now().String(), " +0000 UTC", "", 1)

	updateUser, err := domain.NewUser(user.UserId, user.UserName, user.DisplayName, createdDateTime, updatedDateTime)
	if err != nil {
		return err
	}
	if err := u.userRepo.Save(ctx, updateUser); err != nil {
		return err
	}
	return nil
}

func (u *UserService) DeleteUserById(ctx context.Context, id string) error {
	return u.userRepo.Delete(ctx, id)
}
