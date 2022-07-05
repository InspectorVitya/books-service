package grpcserver

import (
	"books-service/book/pb"
	"books-service/internal/app"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	app        *app.BooksApp
	GRPCServer *grpc.Server
	pb.UnimplementedBookStorageServer
	port string
}

func New(app *app.BooksApp, port string) *Server {
	return &Server{
		app:        app,
		GRPCServer: grpc.NewServer(),
		port:       port,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		return fmt.Errorf("start grpc server failed: %w", err)
	}
	pb.RegisterBookStorageServer(s.GRPCServer, s)
	return s.GRPCServer.Serve(listener)
}

func (s *Server) Stop() {
	s.GRPCServer.GracefulStop()
}
