package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"

	"github.com/ryohei1216/gqlgen-resolver/graph/model"
)

// Books is the resolver for the books field.
func (r *authorResolver) Books(ctx context.Context, obj *model.Author) ([]*model.Book, error) {
	books := make([]*model.Book, 0)
	for _, book := range r.Resolver.Books {
		if book.AuthorID == obj.ID {
			books = append(books, book)
		}
	}

	return books, nil
}

// CreateBook is the resolver for the createBook field.
func (r *mutationResolver) CreateBook(ctx context.Context, title string, authorID int) (*model.Book, error) {
	book := &model.Book{
		ID:       len(r.Books) + 1,
		Title:    title,
		AuthorID: authorID,
	}

	r.Books = append(r.Books, book)

	return book, nil
}

// CreateAuthor is the resolver for the createAuthor field.
func (r *mutationResolver) CreateAuthor(ctx context.Context, name string) (*model.Author, error) {
	author := &model.Author{
		ID:   len(r.Authors) + 1,
		Name: name,
	}

	r.Authors = append(r.Authors, author)

	return author, nil
}

// Author is the resolver for the author field.
func (r *queryResolver) Author(ctx context.Context, id int) (*model.Author, error) {
	var author *model.Author

	for _, a := range r.Authors {
		if a.ID == id {
			author = a
			break
		}
	}

	return author, nil
}

// Author returns AuthorResolver implementation.
func (r *Resolver) Author() AuthorResolver { return &authorResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type authorResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
