package domain

import (
	"strings"

	"github.com/o-ga09/graphql-go/pkg/uuid"
)

/*
Note はノート情報を表す構造体です。
ノート情報は以下の情報を持ちます。
- ID: ノートの ID
- UserID: ユーザーの ID
- Content: ノートの内容
- Title: ノートのタイトル
- Tags: ノートのタグ
- CreatedDateTime: ノートの作成日時
- UpdatedDateTime: ノートの更新日時
*/
type Note struct {
	ID              string
	UserID          string
	Content         string
	Title           string
	Tags            []string
	CreatedDateTime string
	UpdatedDateTime string
}

/*
NewNote はノート情報を生成する関数です。
ノート情報を生成する際には、以下の情報を引数に渡す必要があります。
- userId: ユーザーの ID
- title: ノートのタイトル
- content: ノートの内容
- tags: ノートのタグ
- created: ノートの作成日時
- updated: ノートの更新日時
*/
func NewNote(userId, title, content, tags, created, updated string) (*Note, error) {
	t := strings.Split(tags, ",")
	id := uuid.GenerateID()

	return &Note{
		ID:              id,
		UserID:          userId,
		Content:         content,
		Title:           title,
		Tags:            t,
		CreatedDateTime: created,
		UpdatedDateTime: updated,
	}, nil
}

func ReConstractNote(id, userId, title, content, tags, created, updated string) (*Note, error) {
	t := strings.Split(tags, ",")

	return &Note{
		ID:              id,
		UserID:          userId,
		Content:         content,
		Title:           title,
		Tags:            t,
		CreatedDateTime: created,
		UpdatedDateTime: updated,
	}, nil

}
