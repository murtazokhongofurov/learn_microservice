package grpcclient

import (
	"github.com/three_serviceSimple/store_service/config"
)

type GrpcClientI interface {
}

type ServiceManager struct {
	cfg         config.Config
	connections map[string]interface{}
}

func New(cfg config.Config) (*ServiceManager, error) {
	return &ServiceManager{
		cfg:         cfg,
		connections: map[string]interface{}{},
	}, nil
}
