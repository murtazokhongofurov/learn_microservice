package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/three_serviceSimple/store_service/kafka"
	l "github.com/three_serviceSimple/store_service/pkg/logger"
	"github.com/three_serviceSimple/store_service/storage"
)

type StoreService struct {
	storage  storage.IStorage
	logger   l.Logger
	producer map[string]kafka.Producer
}

func NewStoreService(db *sqlx.DB, log l.Logger, producer map[string]kafka.Producer) *StoreService {
	return &StoreService{
		storage:  storage.NewStoragePg(db),
		logger:   log,
		producer: producer,
	}
}
