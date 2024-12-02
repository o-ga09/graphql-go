package dao

import (
	"context"
	"database/sql"
	"strings"

	"github.com/o-ga09/graphql-go/internal/db/db"
	"github.com/o-ga09/graphql-go/internal/domain"
)

type userDao struct {
	query *db.Queries
}

func NewUserDao(d *sql.DB) *userDao {
	q := db.New(d)
	return &userDao{
		query: q,
	}
}

func (u *userDao) GetUsers(ctx context.Context) ([]*domain.User, error) {
	users, err := u.query.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	res := []*domain.User{}
	for _, user := range users {
		createdDateTime := strings.Replace(user.CreatedAt.Time.String(), " +0000 UTC", "", 1)
		updatedDateTime := strings.Replace(user.UpdatedAt.Time.String(), " +0000 UTC", "", 1)
		r, err := domain.ReconstractUser(user.UserID, user.Username, user.Displayname, createdDateTime, updatedDateTime)
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}

func (u *userDao) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	user, err := u.query.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	createdDateTime := strings.Replace(user.CreatedAt.Time.String(), " +0000 UTC", "", 1)
	updatedDateTime := strings.Replace(user.UpdatedAt.Time.String(), " +0000 UTC", "", 1)
	return domain.ReconstractUser(user.UserID, user.Username, user.Displayname, createdDateTime, updatedDateTime)
}

func (u *userDao) Save(ctx context.Context, user *domain.User) error {
	record, _ := u.query.GetUser(ctx, user.ID)
	if record.UserID != user.ID {
		_, err := u.query.CreateUser(ctx, db.CreateUserParams{
			UserID:      user.ID,
			Username:    user.UserName,
			Displayname: user.DisplayName,
		})
		if err != nil {
			return err
		}
	} else {
		err := u.query.UpdateUser(ctx, db.UpdateUserParams{
			UserID:      user.ID,
			Username:    user.UserName,
			Displayname: user.DisplayName,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *userDao) Delete(ctx context.Context, id string) error {
	err := u.query.DeleteUser(ctx, id)
	return err
}
