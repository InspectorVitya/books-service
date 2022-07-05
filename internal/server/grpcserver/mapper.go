package grpcserver

import (
	"books-service/book/pb"
	"books-service/internal/model"
)

// maps internal model to external
func mapBookListToProtoBookList(books []model.Books) []*pb.Book {

	result := make([]*pb.Book, 0, len(books))
	for _, book := range books {
		result = append(result, &pb.Book{
			Id:          book.ID,
			Title:       book.Title,
			Description: book.Description,
		})
	}
	return result
}

// maps internal model to external
func mapAuthorListToProtoAuthorList(authors []model.Author) []*pb.Author {
	result := make([]*pb.Author, 0, len(authors))
	for _, author := range authors {
		result = append(result, &pb.Author{
			Id:      author.ID,
			Name:    author.Name,
			Surname: author.Surname,
		})
	}
	return result
}
