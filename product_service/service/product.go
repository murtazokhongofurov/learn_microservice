package service
import (
	"context"

	pb "github.com/three_serviceSimple/product_service/genproto/product"
	l "github.com/three_serviceSimple/product_service/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ProductService) GetProductById(ctx context.Context, req *pb.ProductId) (*pb.Product, error) {
	product, err := s.storage.Product().GetProductById(req)
	if err != nil {
		s.logger.Error("error select product", l.Any("Error select product", err))
		return &pb.Product{}, status.Error(codes.Internal, "something went wrong, please check product info")
	}
	return product, nil
}
func (s *ProductService) UpdateProduct(ctx context.Context, req *pb.Product) (*pb.Product, error) {
	product, err := s.storage.Product().UpdateProduct(req)
	if err != nil {
		s.logger.Error("error update product", l.Any("Error update product", err))
		return &pb.Product{}, status.Error(codes.Internal, "something went wrong, please check product info")
	}
	return product, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, req *pb.ProductId) (*pb.Empty, error) {
	product, err := s.storage.Product().DeleteProduct(req)
	if err != nil {
		s.logger.Error("error update product", l.Any("Error update product", err))
		return &pb.Empty{}, status.Error(codes.Internal, "something went wrong, please check product info")
	}
	return product, nil
}
