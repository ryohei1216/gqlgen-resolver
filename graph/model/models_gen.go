// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Author struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Books []*Book `json:"books"`
}

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	AuthorID int    `json:"authorId"`
}
