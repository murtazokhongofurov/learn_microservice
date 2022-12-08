package grpcclient

import (
	"fmt"

	"github.com/three_serviceSimple/product_service/config"
	pbp "github.com/three_serviceSimple/product_service/genproto/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServceManager struct {
	cfg          config.Config
	storeService pbp.StoreServiceClient
}

func New(conf config.Config) (*ServceManager, error) {
	connStore, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.StoreServiceHost, conf.StoreServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("error while dial customer service: host: %s and port: %d",
			conf.StoreServiceHost, conf.StoreServicePort)
	}
	serviceManager := &ServceManager{
		cfg:          conf,
		storeService: pbp.NewStoreServiceClient(connStore),
	}
	return serviceManager, nil
}
func (s ServceManager) StoreService() pbp.StoreServiceClient {
	return s.storeService
}

// type GrpcClientI interface {
// }

// type GrpcClient struct {
// 	cfg         config.Config
// 	connections map[string]interface{}
// }

// func New(cfg config.Config) (*GrpcClient, error) {
// 	return &GrpcClient{
// 		cfg:         cfg,
// 		connections: map[string]interface{}{},
// 	}, nil
// }
