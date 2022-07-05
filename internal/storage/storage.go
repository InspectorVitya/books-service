package storage

import (
	"books-service/internal/model"
	"context"
)

type BooksStorage interface {
	FindBooksByAuthor(ctx context.Context, id uint64) ([]model.Books, error)
	FindAuthorByBook(ctx context.Context, id uint64) ([]model.Author, error)
}
