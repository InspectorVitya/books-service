package app

import (
	"books-service/internal/model"
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

type memoryDB struct{}

func (s *memoryDB) FindBooksByAuthor(ctx context.Context, id uint64) ([]model.Books, error) {
	books := make([]model.Books, 0)
	switch id {
	case 1:
		books = append(books, model.Books{
			ID:          0,
			Title:       "Harry Potter",
			Description: "classic book",
		})
	case 2:
		books = append(books, model.Books{
			Title:       "Anna karenina",
			Description: "classic book",
		})
		books = append(books, model.Books{
			Title:       "War and Peace",
			Description: "classic book",
		})
	case 3:
		return nil, errors.New("errors")
	}

	return books, nil
}

func (s *memoryDB) FindAuthorByBook(ctx context.Context, id uint64) ([]model.Author, error) {
	authors := make([]model.Author, 0)

	switch id {
	case 1:
		authors = append(authors, model.Author{
			Name:    "Lev",
			Surname: "Tolstoy",
		})
	case 2:
		authors = append(authors, model.Author{
			Name:    "Lev",
			Surname: "Tolstoy",
		})
		authors = append(authors, model.Author{
			Name:    "J",
			Surname: "Rowling",
		})
	case 3:
		return nil, errors.New("errors")
	}
	return authors, nil
}

func Test_GetBooks(t *testing.T) {
	db := memoryDB{}
	bookApp := New(&db)
	t.Run("one result", func(t *testing.T) {
		data, err := bookApp.GetBooks(context.Background(), 1)
		require.NoError(t, err)
		require.Equal(t, 1, len(data))
	})
	t.Run("two result", func(t *testing.T) {
		data, err := bookApp.GetBooks(context.Background(), 2)
		require.NoError(t, err)
		require.Equal(t, 2, len(data))
	})
	t.Run("with error", func(t *testing.T) {
		data, err := bookApp.GetBooks(context.Background(), 3)
		require.Error(t, err)
		require.Empty(t, data)
	})
	t.Run("with empty result", func(t *testing.T) {
		data, err := bookApp.GetBooks(context.Background(), 4)
		require.NoError(t, err)
		require.Empty(t, data)
	})
}

func Test_GetAuthor(t *testing.T) {
	db := memoryDB{}
	bookApp := New(&db)
	t.Run("one result", func(t *testing.T) {
		data, err := bookApp.GetAuthor(context.Background(), 1)
		require.NoError(t, err)
		require.Equal(t, 1, len(data))
	})
	t.Run("two result", func(t *testing.T) {
		data, err := bookApp.GetAuthor(context.Background(), 2)
		require.NoError(t, err)
		require.Equal(t, 2, len(data))
	})
	t.Run("with error", func(t *testing.T) {
		data, err := bookApp.GetAuthor(context.Background(), 3)
		require.Error(t, err)
		require.Empty(t, data)
	})
	t.Run("with empty result", func(t *testing.T) {
		data, err := bookApp.GetAuthor(context.Background(), 4)
		require.NoError(t, err)
		require.Empty(t, data)
	})
}
