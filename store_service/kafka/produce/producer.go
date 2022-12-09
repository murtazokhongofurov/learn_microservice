package producer

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/three_serviceSimple/store_service/config"
	producer "github.com/three_serviceSimple/store_service/kafka"
	"github.com/three_serviceSimple/store_service/pkg/logger"
)

type KafkaProduce struct {
	kafkaWriter *kafka.Writer
	log         logger.Logger
}

func NewKafkaProducer(cfg config.Config, log logger.Logger, topic string) producer.Producer {
	return &KafkaProduce{
		kafkaWriter: &kafka.Writer{
			Addr:         kafka.TCP(cfg.KafkaHost),
			Topic:        topic,
			BatchTimeout: time.Second * 2,
		},
		log: log,
	}
}

func (p *KafkaProduce) Start() error {
	return nil
}

func (p *KafkaProduce) Stop() error {
	err := p.kafkaWriter.Close()
	if err != nil {
		return err
	}
	return nil
}

func (p KafkaProduce) Produce(key, body []byte, logBody string) error {
	message := kafka.Message{
		Key:   key,
		Value: body,
	}
	if err := p.kafkaWriter.WriteMessages(context.Background(), message); err != nil {
		return err
	}
	return nil
}

// type Product struct {
// 	Conn      *kafka.Conn
// 	ConnClose func()
// }

// func NewProduct(cfg config.Config) (*Product, error) {
// 	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "my_topic", 0)
// 	if err != nil {
// 		return &Product{}, err
// 	}
// 	return &Product{
// 		Conn: conn,
// 		ConnClose: func() {
// 			conn.Close()
// 		},
// 	}, err
// }

// func (p *Product) CreateProduct(req *store.ProductRequest) error {
// 	defer p.ConnClose()
// 	byteReq, err := json.Marshal(req)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(byteReq)
// 	_, err = p.Conn.Write(byteReq)
// 	return err
// }
