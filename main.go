package main

import (
	"context"
	pb "github.com/seftomsk/grpc-auth-server/proto/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"strconv"
)

type helloService struct {}

func (s *helloService) Hello(_ context.Context, req *pb.Request) (*pb.Response, error) {
	age := strconv.FormatInt(int64(req.Age), 10)
	return &pb.Response{Message: "Hello " + req.Name + " " + age}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
	if err != nil {
		log.Fatalf("failed to check the certificate: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterHelloServiceServer(grpcServer, &helloService{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
