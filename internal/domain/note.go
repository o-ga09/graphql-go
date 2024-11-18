package domain

import (
	"strings"
	"time"

	"github.com/o-ga09/graphql-go/pkg/date"
)

type Note struct {
	ID              string
	UserID          string
	Content         string
	Title           string
	Tags            []string
	CreatedDateTime time.Time
	UpdatedDateTime time.Time
}

func NewNote(id, userId, title, content, tags, created, updated string) (*Note, error) {
	t := strings.Split(tags, ",")
	CreatedDateTime, err := date.TimeToString(created)
	if err != nil {
		return nil, err
	}
	UpdatedDateTime, err := date.TimeToString(updated)
	if err != nil {
		return nil, err
	}
	return &Note{
		ID:              id,
		UserID:          userId,
		Content:         content,
		Title:           title,
		Tags:            t,
		CreatedDateTime: CreatedDateTime,
		UpdatedDateTime: UpdatedDateTime,
	}, nil
}
