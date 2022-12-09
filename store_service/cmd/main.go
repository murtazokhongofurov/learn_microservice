package main

import (
	"net"

	"github.com/three_serviceSimple/store_service/config"
	"github.com/three_serviceSimple/store_service/kafka"
	produce "github.com/three_serviceSimple/store_service/kafka/produce"
	"github.com/three_serviceSimple/store_service/pkg/db"
	"github.com/three_serviceSimple/store_service/pkg/logger"
	"github.com/three_serviceSimple/store_service/service"

	pb "github.com/three_serviceSimple/store_service/genproto/store"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "golang")
	defer logger.Cleanup(log)
	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase),
	)
	connDB, err := db.ConnectToDB(cfg)

	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}
	produceMap := make(map[string]kafka.Producer)
	topic := "store.store"
	storeTopicProduce := produce.NewKafkaProducer(cfg, log, topic)
	defer func() {
		err := storeTopicProduce.Stop()
		if err != nil {
			log.Fatal("Failed to stopping Kafka", logger.Error(err))
		}
	}()
	produceMap["store"] = storeTopicProduce
	storeService := service.NewStoreService(connDB, log, produceMap)

	lis, err := net.Listen("tcp", ":"+cfg.StoreServicePort)

	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterStoreServiceServer(s, storeService)

	log.Info("main: server running",
		logger.String("port", cfg.StoreServicePort))
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
