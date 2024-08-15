package dto

type NoteDto struct {
	ID              string   `json:"id"`
	Title           string   `json:"title"`
	Content         string   `json:"content"`
	Tags            []string `json:"tags"`
	User            UserDto  `json:"user"`
	CreatedDateTime string   `json:"createdDateTime"`
	UpdatedDateTime string   `json:"updatedDateTime"`
}

type UserDto struct {
	ID              string `json:"id"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Email           string `json:"email"`
	Address         string `json:"address"`
	BirthDay        string `json:"birthDay"`
	Sex             string `json:"sex"`
	Password        string `json:"password"`
	CreatedDateTime string `json:"createdDateTime"`
	UpdatedDateTime string `json:"updatedDateTime"`
}
