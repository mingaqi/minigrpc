package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"minigrpc/pb"
)

func main() {
	s, err := grpc.Dial(":5051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	client := pb.NewProdServiceClient(s)
	prodResponse, err := client.GetProdStock(context.Background(), &pb.ProdRequest{
		ProdId: 4,
	})
	log.Println(prodResponse.GetProdStock())
}
