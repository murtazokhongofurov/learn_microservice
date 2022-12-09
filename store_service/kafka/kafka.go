package kafka

type Producer interface {
	Start() error
	Stop() error
	Produce(key, body []byte, logBody string) error
}

// import (
// 	"gitlab.com/go_kafka/store_service/config"
// 	producer "gitlab.com/go_kafka/store_service/kafka/producer/producer"
// )

// type Kafka struct {
// 	Products *producer.Product
// }

// type KafkaI interface {
// 	Product() *producer.Product
// }

// func NewKafka(cfg config.Config) (KafkaI, func(), error) {
// 	temp, err := producer.NewProduct(cfg)
// 	if err != nil {
// 		return &Kafka{}, func() {}, err
// 	}
// 	return &Kafka{
// 			Products: temp,
// 		}, func() {
// 			temp.ConnClose()
// 		}, nil
// }

// func (k *Kafka) Product() *producer.Product {
// 	return k.Products
// }
