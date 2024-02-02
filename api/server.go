package api

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/Mkukarev/go-grpc/config"
	pb "github.com/Mkukarev/go-grpc/proto"
	"github.com/Mkukarev/go-grpc/store"
)

type Server struct {
	cfg   config.HTTPServer
	store store.Interface
	pb.UnimplementedTodoCRUDServer
}

func NewServer(cfg config.HTTPServer, store store.Interface) *Server {
	s := &Server{
		cfg:   cfg,
		store: store,
	}

	return s
}

func (server *Server) Start(ctx context.Context) {
	// listen, err := net.Listen("tcp", fmt.Sprintf(":%d", server.cfg.Port))
	fmt.Println(server.cfg)
	listen, err := net.Listen("tcp", ":3002")
	if err != nil {
		fmt.Printf("Could not listen on port: %v", err)
	}

	// gRPC server
	s := grpc.NewServer()
	pb.RegisterTodoCRUDServer(s, server)

	fmt.Printf("Hosting server on: %s", listen.Addr().String())

	if err := s.Serve(listen); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}

}
