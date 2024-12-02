package dto

type NoteDto struct {
	ID              string   `json:"id"`
	Title           string   `json:"title"`
	Content         string   `json:"content"`
	Tags            []string `json:"tags"`
	UserId          string   `json:"user"`
	CreatedDateTime string   `json:"createdDateTime"`
	UpdatedDateTime string   `json:"updatedDateTime"`
}

type NoteRequestDto struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

type UserDto struct {
	ID              string `json:"id"`
	UserName        string `json:"userName"`
	DisplayName     string `json:"displayName"`
	CreatedDateTime string `json:"createdDateTime"`
	UpdatedDateTime string `json:"updatedDateTime"`
}

type UserReqsutDto struct {
	UserId      string `json:"userID"`
	UserName    string `json:"userName"`
	DisplayName string `json:"displayName"`
}
