package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Envirnment         string
	PostgresHost       string
	PostgresPort       int
	PostgresDatabase   string
	PostgresUser       string
	PostgresPassword   string
	LogLevel           string
	ProductServicePort string
	ProductServiceHost string
	KafkaHost          string
	StoreServiceHost   string
	StoreServicePort   int
}

func Load() Config {
	c := Config{}
	c.Envirnment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "product_service"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "developer"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "2002"))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.ProductServicePort = cast.ToString(getOrReturnDefault("PRODUCT_SERVICE_PORT", "5000"))
	c.ProductServiceHost = cast.ToString(getOrReturnDefault("PRODUCT_SERVICE_HOST", "product"))
	c.KafkaHost = cast.ToString(getOrReturnDefault("KAFKA_HOST", "kafka:9092"))
	c.StoreServiceHost = cast.ToString(getOrReturnDefault("STORE_SERVICE_HOST", "store"))
	c.StoreServicePort = cast.ToInt(getOrReturnDefault("STORE_SERVICE_PORT", 3000))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}