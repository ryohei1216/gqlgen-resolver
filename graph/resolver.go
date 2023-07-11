package graph

import "github.com/ryohei1216/gqlgen-resolver/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Authors []*model.Author
	Books   []*model.Book
}
