package mysql

import (
	"books-service/internal/model"
	"context"
)

func (s *StorageMySQL) FindBooksByAuthor(ctx context.Context, id uint64) ([]model.Books, error) {
	query := `SELECT title, description FROM  books
		JOIN author_book ON books.book_id = author_book.book_id 
		WHERE author_book.author_id  = ?;`

	var books []model.Books
	err := s.db.SelectContext(ctx, &books, query, id)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (s *StorageMySQL) FindAuthorByBooks(ctx context.Context, id uint64) ([]model.Author, error) {
	query := `SELECT name, surname FROM  author
	JOIN author_book ON author.author_id  = author_book.author_id 
	WHERE author_book.author_id  = ?;`

	var authors []model.Author
	err := s.db.SelectContext(ctx, &authors, query, id)
	if err != nil {
		return nil, err
	}
	return authors, nil
}
