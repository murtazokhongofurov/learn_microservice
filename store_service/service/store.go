package service
import (
	"context"
	"fmt"

	pb "github.com/three_serviceSimple/store_service/genproto/store"
	l "github.com/three_serviceSimple/store_service/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *StoreService) produceMessage(raw *pb.ProductRes) error {
	data, err := raw.Marshal()
	if err != nil {
		return err
	}
	logPost := raw.String()
	err = s.producer["store"].Produce([]byte("store"), data, logPost)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

func (s *StoreService) CreateStore(ctx context.Context, req *pb.StoreRequest) (*pb.StoreResponse, error) {
	store, err := s.storage.Store().CreateStore(req)
	if err != nil {
		s.logger.Error("error insert store", l.Any("Error insert store", err))
		return &pb.StoreResponse{}, status.Error(codes.Internal, "something went wrong, please check store info")
	}
	err = s.produceMessage(store.Products[codes.Aborted])
	if err != nil {
		fmt.Println("error while produce products")
		return &pb.StoreResponse{}, err
	}
	return store, nil
}
func (s *StoreService) GetStoreInfo(ctx context.Context, req *pb.StoreId) (*pb.StoreResponse, error) {
	store, err := s.storage.Store().GetStoreInfo(req)
	if err != nil {
		s.logger.Error("error select store", l.Any("error select store", err))
		return &pb.StoreResponse{}, status.Error(codes.Internal, "something went wrong, please check store info")
	}
	return store, nil
}
func (s *StoreService) UpdateStore(ctx context.Context, req *pb.StoreResponse) (*pb.StoreResponse, error) {
	store, err := s.storage.Store().UpdateStore(req)
	if err != nil {
		s.logger.Error("error update store", l.Any("error update store", err))
		return &pb.StoreResponse{}, status.Error(codes.Internal, "something went wrong, please check store info")
	}
	return store, nil
}
func (s *StoreService) DeleteStore(ctx context.Context, req *pb.StoreId) (*pb.Empty, error) {
	store, err := s.storage.Store().DeleteStore(req)
	if err != nil {
		s.logger.Error("error delete store", l.Any("error delete store", err))
		return &pb.Empty{}, status.Error(codes.Internal, "something went wrong, please check store info")
	}
	return store, nil
}
