package main

import (
	"context"
	"log"
	"net"

	protocol "github.com/Meheer17/grpc-golang-demo/protobuf/protocol_demo"
	"google.golang.org/grpc"
)

type MyProtocolServer struct {
	protocol.UnimplementedProtocolServer
}

func (s MyProtocolServer) Create(ctx context.Context, req  *protocol.CreateRequest) (*protocol.CreateResponse, error) {
	return &protocol.CreateResponse{
		Id: req.From + " " + req.To,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	
	if err != nil {
		log.Fatalf("Can't listen on port %s", ":8080")
	}
	
	serverRegistar := grpc.NewServer()
	service := &MyProtocolServer{}
	
	protocol.RegisterProtocolServer(serverRegistar, service)
	err = serverRegistar.Serve(lis)
	
	if err != nil {
		log.Fatalf("Impossible to serve gRPC server on port %s", ":8080")
	}
}