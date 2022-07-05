package grpcserver

import (
	"books-service/book/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetBooksByAuthorId(ctx context.Context, req *pb.BookRequest) (*pb.BookList, error) {
	id := req.AuthorId
	if id == 0 {
		return nil, status.Error(codes.InvalidArgument, "author id is not specified")
	}
	books, err := s.app.GetBooks(ctx, id)
	if err != nil {
		return nil, err
	}
	response := &pb.BookList{}
	response.Book = mapBookListToProtoBookList(books)
	return response, nil
}

func (s *Server) GetAuthorsByBookId(ctx context.Context, req *pb.AuthorRequest) (*pb.AuthorList, error) {
	id := req.BookId
	if id == 0 {
		return nil, status.Error(codes.InvalidArgument, "book id is not specified")
	}
	authors, err := s.app.GetAuthor(ctx, id)
	if err != nil {
		return nil, err
	}
	response := &pb.AuthorList{}
	response.Author = mapAuthorListToProtoAuthorList(authors)

	return response, nil
}
