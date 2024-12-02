package domain

import (
	"errors"
)

var (
	ErrDateFormat = errors.New("date format is invalid")
)

/*
User はユーザー情報を表す構造体です。
ユーザー情報は以下の情報を持ちます。
- ID: firebase の uid
- UserName: ユーザー名
- DisplayName: ユーザーの表示名
- CreatedDateTime: ユーザーの作成日時
- UpdatedDateTime: ユーザーの更新日時
*/
type User struct {
	ID              string
	UserName        string
	DisplayName     string
	CreatedDateTime string
	UpdatedDateTime string
}

/*
NewUser はユーザー情報を生成する関数です。
ユーザー情報を生成する際には、以下の情報を引数に渡す必要があります。
- username: ユーザー名
- displayname: ユーザーの表示名
*/
func NewUser(id, username, displayname, createdDateTime, updatedDateTime string) (*User, error) {
	return &User{
		ID:              id,
		UserName:        username,
		DisplayName:     displayname,
		CreatedDateTime: createdDateTime,
		UpdatedDateTime: updatedDateTime,
	}, nil
}

func ReconstractUser(id, username, displayname, created, updated string) (*User, error) {
	return &User{
		ID:              id,
		UserName:        username,
		DisplayName:     displayname,
		CreatedDateTime: created,
		UpdatedDateTime: updated,
	}, nil
}
