package grpcclient

import (
	"github.com/three_serviceSimple/user_service/config"
)

type GrpcClientI interface {
}

type GrpcClient struct {
	cfg config.Config
}

func New(cfg config.Config) (*GrpcClient, error) {

	return &GrpcClient{
		cfg: cfg,
	}, nil
}
