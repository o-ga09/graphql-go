// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/google/uuid"
)

type DeleteNote struct {
	ID string `json:"id"`
}

type DeleteUser struct {
	ID string `json:"id"`
}

type DeletedNote struct {
	ID string `json:"id"`
}

type DeletedUser struct {
	ID string `json:"id"`
}

type EditedNote struct {
	ID string `json:"id"`
}

type EditedUser struct {
	ID string `json:"id"`
}

type Mutation struct {
}

type NewNote struct {
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
	Title   string   `json:"title"`
}

type NewUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Sex       int    `json:"sex"`
	Password  string `json:"password"`
	BirthDay  string `json:"birth_day"`
}

type Note struct {
	ID              uuid.UUID `json:"id"`
	Content         string    `json:"content"`
	Tags            []string  `json:"tags"`
	Title           string    `json:"title"`
	User            *User     `json:"user"`
	CreatedDateTime time.Time `json:"CreatedDateTime"`
	UpdatedDateTime time.Time `json:"UpdatedDateTime"`
}

type Query struct {
}

type UpdateNote struct {
	ID      string    `json:"id"`
	Content *string   `json:"content,omitempty"`
	Tags    []*string `json:"tags,omitempty"`
	Title   *string   `json:"title,omitempty"`
}

type UpdateUser struct {
	ID        string  `json:"id"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Email     *string `json:"email,omitempty"`
	Address   *string `json:"address,omitempty"`
	Sex       *int    `json:"sex,omitempty"`
	Password  *string `json:"password,omitempty"`
	BirthDay  *string `json:"birth_day,omitempty"`
}

type User struct {
	ID              uuid.UUID `json:"id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Email           string    `json:"email"`
	Address         string    `json:"address"`
	Sex             int       `json:"sex"`
	Password        string    `json:"password"`
	BirthDay        string    `json:"birth_day"`
	CreatedDateTime time.Time `json:"CreatedDateTime"`
	UpdatedDateTime time.Time `json:"UpdatedDateTime"`
}
