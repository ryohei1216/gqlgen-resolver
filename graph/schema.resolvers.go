package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"fmt"

	"github.com/ryohei1216/gqlgen-resolver/graph/model"
)

// CreateBook is the resolver for the createBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, title string, authorID int) (*model.Book, error) {
	panic(fmt.Errorf("not implemented: CreateBook - createBook"))
}

// CreateAuthor is the resolver for the createAuthor field.
func (r *mutationResolver) CreateAuthor(ctx context.Context, name string) (*model.Author, error) {
	panic(fmt.Errorf("not implemented: CreateAuthor - createAuthor"))
}

// Author is the resolver for the author field.
func (r *queryResolver) Author(ctx context.Context, id int) (*model.Author, error) {
	panic(fmt.Errorf("not implemented: Author - author"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
