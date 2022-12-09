package messagebroker

type Producer interface {
	Start() error
	Stop() error
	Publish(key, body []byte, logBody string) error
}