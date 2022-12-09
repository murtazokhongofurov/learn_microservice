package repo

import (
	pb "github.com/three_serviceSimple/store_service/genproto/store"
)

type StoreStorageI interface {
	CreateStore(*pb.StoreRequest) (*pb.StoreResponse, error)
	GetStoreInfo(*pb.StoreId) (*pb.StoreResponse, error)
	UpdateStore(*pb.StoreResponse) (*pb.StoreResponse, error)
	DeleteStore(*pb.StoreId) (*pb.Empty, error)
}
