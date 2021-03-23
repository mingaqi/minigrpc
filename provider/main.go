package main

import (
	"google.golang.org/grpc"
	"log"
	"minigrpc/pb"
	"minigrpc/provider/services"
	"net"
)

func main() {

	lis, err := net.Listen("tcp", ":5051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProdServiceServer(s, &services.ProdService{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
