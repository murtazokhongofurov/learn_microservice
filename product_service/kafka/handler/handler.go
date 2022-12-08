package handler

import (
	"github.com/three_serviceSimple/product_service/config"
	pb "github.com/three_serviceSimple/product_service/genproto/store"
	"github.com/three_serviceSimple/product_service/pkg/logger"
	"github.com/three_serviceSimple/product_service/storage"
)

type KafkaHandler struct {
	config  config.Config
	storage storage.IStorage
	log     logger.Logger
}

func NewKafkaHandlerFunc(config config.Config, storage storage.IStorage, log logger.Logger) *KafkaHandler {
	return &KafkaHandler{
		storage: storage,
		config:  config,
		log:     log,
	}
}

func (h *KafkaHandler) Handle(value []byte) error {
	product := pb.ProductRes{}
	err := product.Unmarshal(value)
	if err != nil {
		return err
	}
	err = h.storage.Product().CreateProduct(&product)
	if err != nil {
		return err
	}
	return nil
}

// type ProductConsumer struct {
// 	Reader    *kafka.Reader
// 	ConnClose func()
// 	Cfg       config.Config
// }

// func NewProductConcumer(cfg config.Config) (*ProductConsumer, error) {
// 	r := kafka.NewReader(kafka.ReaderConfig{
// 		Brokers:   []string{"localhost:9092", "localhost:9093", "localhost:9094", "localhost:29092", "localhost:7000"},
// 		Topic:     cfg.ProductTopic,
// 		Partition: 0,
// 		MinBytes:  1e3,  //10 KB
// 		MaxBytes:  10e6, //10 MB
// 	})
// 	return &ProductConsumer{
// 		Reader: r,
// 		ConnClose: func() {
// 			r.Close()
// 		},
// 	}, nil
// }

// func (p *ProductConsumer) Product() (*product.Empty, error) {
// 	for {
// 		fmt.Println("going")
// 		req := &product.Product{}

// 		m, err := p.Reader.ReadMessage(context.Background())
// 		if err != nil {
// 			break
// 		}
// 		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
// 		err = json.Unmarshal(m.Value, &req)
// 		if err != nil {
// 			fmt.Println(err)
// 			break
// 		}
// 	}
// 	return &product.Empty{}, nil
// }
