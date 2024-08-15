package graph

import "github.com/o-ga09/graphql-go/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	NoteService service.NoteService
	UserService service.UserService
}
