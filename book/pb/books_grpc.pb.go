// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BookStorageClient is the client API for BookStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookStorageClient interface {
	GetBooksByAuthorId(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookList, error)
	GetAuthorsByBookId(ctx context.Context, in *AuthorRequest, opts ...grpc.CallOption) (*AuthorList, error)
}

type bookStorageClient struct {
	cc grpc.ClientConnInterface
}

func NewBookStorageClient(cc grpc.ClientConnInterface) BookStorageClient {
	return &bookStorageClient{cc}
}

func (c *bookStorageClient) GetBooksByAuthorId(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookList, error) {
	out := new(BookList)
	err := c.cc.Invoke(ctx, "/book.BookStorage/GetBooksByAuthorId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookStorageClient) GetAuthorsByBookId(ctx context.Context, in *AuthorRequest, opts ...grpc.CallOption) (*AuthorList, error) {
	out := new(AuthorList)
	err := c.cc.Invoke(ctx, "/book.BookStorage/GetAuthorsByBookId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookStorageServer is the server API for BookStorage service.
// All implementations must embed UnimplementedBookStorageServer
// for forward compatibility
type BookStorageServer interface {
	GetBooksByAuthorId(context.Context, *BookRequest) (*BookList, error)
	GetAuthorsByBookId(context.Context, *AuthorRequest) (*AuthorList, error)
	mustEmbedUnimplementedBookStorageServer()
}

// UnimplementedBookStorageServer must be embedded to have forward compatible implementations.
type UnimplementedBookStorageServer struct {
}

func (UnimplementedBookStorageServer) GetBooksByAuthorId(context.Context, *BookRequest) (*BookList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooksByAuthorId not implemented")
}
func (UnimplementedBookStorageServer) GetAuthorsByBookId(context.Context, *AuthorRequest) (*AuthorList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorsByBookId not implemented")
}
func (UnimplementedBookStorageServer) mustEmbedUnimplementedBookStorageServer() {}

// UnsafeBookStorageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookStorageServer will
// result in compilation errors.
type UnsafeBookStorageServer interface {
	mustEmbedUnimplementedBookStorageServer()
}

func RegisterBookStorageServer(s grpc.ServiceRegistrar, srv BookStorageServer) {
	s.RegisterService(&BookStorage_ServiceDesc, srv)
}

func _BookStorage_GetBooksByAuthorId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookStorageServer).GetBooksByAuthorId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookStorage/GetBooksByAuthorId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookStorageServer).GetBooksByAuthorId(ctx, req.(*BookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookStorage_GetAuthorsByBookId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookStorageServer).GetAuthorsByBookId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookStorage/GetAuthorsByBookId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookStorageServer).GetAuthorsByBookId(ctx, req.(*AuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BookStorage_ServiceDesc is the grpc.ServiceDesc for BookStorage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookStorage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book.BookStorage",
	HandlerType: (*BookStorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBooksByAuthorId",
			Handler:    _BookStorage_GetBooksByAuthorId_Handler,
		},
		{
			MethodName: "GetAuthorsByBookId",
			Handler:    _BookStorage_GetAuthorsByBookId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "books.proto",
}