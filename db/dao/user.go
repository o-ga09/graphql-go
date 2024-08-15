package dao

import (
	"context"
	"database/sql"

	"github.com/o-ga09/graphql-go/db/db"
	"github.com/o-ga09/graphql-go/domain"
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
		r, err := domain.NewUser(user.UserID, user.Name, user.Email, user.Address, user.Birthday, user.Password, user.CreatedAt.Time.Format("2006-01-02 15:04:05"), user.UpdatedAt.Time.Format("2006-01-02 15:04:05"), int(user.Sex))
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}
	return res, nil
}

func (u *userDao) GetUserById(ctx context.Context, id string) (*domain.User, error) {
	user, err := u.query.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return domain.NewUser(user.UserID, user.Name, user.Email, user.Address, user.Birthday, user.Password, user.CreatedAt.Time.Format("2006-01-02 15:04:05"), user.UpdatedAt.Time.Format("2006-01-02 15:04:05"), int(user.Sex))
}

func (u *userDao) CreateUser(ctx context.Context, user *domain.User) error {
	_, err := u.query.CreateUser(ctx, db.CreateUserParams{
		UserID:   user.ID,
		Name:     user.FirstName + " " + user.LastName,
		Email:    user.Email,
		Address:  user.Address,
		Birthday: user.BirthDay,
		Password: user.Password,
		Sex:      int32(user.Sex),
	})
	return err
}

func (u *userDao) UpdateUserById(ctx context.Context, id string, user *domain.User) error {
	err := u.query.UpdateUser(ctx, db.UpdateUserParams{
		UserID:   user.ID,
		Name:     user.FirstName + " " + user.LastName,
		Email:    user.Email,
		Address:  user.Address,
		Birthday: user.BirthDay,
		Password: user.Address,
	})
	return err
}

func (u *userDao) DeleteUserById(ctx context.Context, id string) error {
	err := u.query.DeleteUser(ctx, id)
	return err
}
