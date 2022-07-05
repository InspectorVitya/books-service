package mysql

import (
	"books-service/internal/model"
	"context"
	"fmt"
)

func (s *StorageMySQL) FindBooksByAuthor(ctx context.Context, id uint64) ([]model.Books, error) {
	query := `SELECT  id, title, description FROM  books
		JOIN author_book ON books.id = author_book.book_id 
		WHERE author_book.author_id  = ?;`

	var books []model.Books
	err := s.db.SelectContext(ctx, &books, query, id)
	if err != nil {
		fmt.Println("sdas")
		return nil, err
	}

	return books, nil
}

func (s *StorageMySQL) FindAuthorByBook(ctx context.Context, id uint64) ([]model.Author, error) {
	query := `SELECT id, name, surname FROM  author
	JOIN author_book ON author.id  = author_book.author_id 
	WHERE author_book.book_id  = ?;`

	var authors []model.Author
	err := s.db.SelectContext(ctx, &authors, query, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return authors, nil
}
