package services

import (
	"context"
	"minigrpc/pb"
)

type ProdService struct {
	pb.UnimplementedProdServiceServer
}

func (s *ProdService) GetProdStock(ctx context.Context, request *pb.ProdRequest) (*pb.ProdResponse, error) {
	return &pb.ProdResponse{ProdStock: request.ProdId * 20}, nil
}
