package server

import "books-service/internal/server/grpcserver"

type Server struct {
	GRPC *grpcserver.Server
}

func New(grpcServer *grpcserver.Server) *Server {
	return &Server{
		GRPC: grpcServer,
	}
}
