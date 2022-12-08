package service
import (
	"context"

	pbp "github.com/three_serviceSimple/product_service/genproto/store"
	l "github.com/three_serviceSimple/product_service/pkg/logger"
	grpcclient "github.com/three_serviceSimple/product_service/service/grpc_client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StoreService struct {
	store  *grpcclient.ServceManager
	logger l.Logger
}

func (s *StoreService) CreateStore(ctx context.Context, req *pbp.StoreRequest) (*pbp.StoreResponse, error) {
	store, err := s.store.StoreService().CreateStore(ctx, req)
	if err != nil {
		s.logger.Error("Error while insert", l.Any("error insert store", err))
		return &pbp.StoreResponse{}, status.Error(codes.Internal, "something went wrong, please check store info")
	}
	return store, nil
}
