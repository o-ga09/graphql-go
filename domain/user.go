package domain

import (
	"time"

	"github.com/o-ga09/graphql-go/pkg/date"
)

type User struct {
	ID              string
	FirstName       string
	LastName        string
	Email           string
	Address         string
	BirthDay        string
	Sex             int
	Password        string
	CreatedDateTime time.Time
	UpdatedDateTime time.Time
}

func NewUser(id, f_name, l_name, email, address, birthday, password, created, updated string, sex int) (*User, error) {
	createdDateTime, err := date.TimeToString(created)
	if err != nil {
		return nil, err
	}
	updatedDateTime, err := date.TimeToString(updated)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:              id,
		FirstName:       f_name,
		LastName:        l_name,
		Email:           email,
		Address:         address,
		BirthDay:        birthday,
		Password:        password,
		Sex:             sex,
		CreatedDateTime: createdDateTime,
		UpdatedDateTime: updatedDateTime,
	}, nil
}
