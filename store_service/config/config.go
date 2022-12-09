package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Envirnment       string
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
	LogLevel         string
	StoreServicePort string
	StoreServiceHost string
	KafkaHost        string
	Partition        int
}

func Load() Config {
	c := Config{}
	c.Envirnment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "store_service"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "developer"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "2002"))
	c.KafkaHost = cast.ToString(getOrReturnDefault("KAFKA_HOST", "kafka:9092"))
	c.Partition = cast.ToInt(getOrReturnDefault("PARTITIONS", 0))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.StoreServicePort = cast.ToString(getOrReturnDefault("STORE_SERVICE_PORT", "3000"))
	c.StoreServiceHost = cast.ToString(getOrReturnDefault("STORE_SERVICE_HOST", "store"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
