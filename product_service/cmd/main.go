package main

import (
	"github.com/three_serviceSimple/product_service/service"
	"net"

	"github.com/three_serviceSimple/product_service/config"
	"github.com/three_serviceSimple/product_service/kafka"
	"github.com/three_serviceSimple/product_service/pkg/db"
	"github.com/three_serviceSimple/product_service/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "github.com/three_serviceSimple/product_service/genproto/product"
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
	StoreCreateTopic := kafka.NewKafkaConsumer(connDB, &cfg, log, "store.store")
	go StoreCreateTopic.Start()

	productService := service.NewProductService(connDB, log)

	lis, err := net.Listen("tcp", ":"+cfg.ProductServicePort)

	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterProductServiceServer(s, productService)

	log.Info("main: server running",
		logger.String("port", cfg.ProductServicePort))
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
