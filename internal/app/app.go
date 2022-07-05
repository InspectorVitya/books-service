package app

import (
	"books-service/internal/model"
	"books-service/internal/storage"
	"context"
)

type BooksApp struct {
	db storage.BooksStorage
}

//New возвращает новый BooksApp
func New(db storage.BooksStorage) *BooksApp {
	return &BooksApp{db: db}
}

//GetBooks возвращает слайс книг по автору и ошибку, если она есть
func (app *BooksApp) GetBooks(ctx context.Context, id uint64) ([]model.Books, error) {
	books, err := app.db.FindBooksByAuthor(ctx, id)
	if err != nil {
		return nil, err
	}
	return books, nil
}

//GetAuthor возвращает слайс авторов по книге и ошибку, если она есть
func (app *BooksApp) GetAuthor(ctx context.Context, id uint64) ([]model.Author, error) {
	authors, err := app.db.FindAuthorByBook(ctx, id)
	if err != nil {
		return nil, err
	}
	return authors, nil
}
