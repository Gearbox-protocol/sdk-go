package log

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type ampqService struct {
	ch *amqp.Channel
}

func failOnError(err error, msg string) {
	if err != nil {
		Fatalf("%s: %s", msg, err)
	}
}

// Service constructor
func NewAMQPService(amqpEnable string, amqpUrl string, logConfig LoggingConfig) {
	if amqpEnable == "0" {
		return
	}
	//
	conn, err := amqp.Dial(amqpUrl)
	failOnError(err, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	//
	logConfig.Channel = ch
	_logConfig = logConfig
}

func GetNetworkName(chainId int64) (name string) {
	switch chainId {
	case 42:
		name = "KOVAN"
	case 5:
		name = "GOERLI"
	case 1:
		name = "MAINNET"
	case 1337:
		name = "TEST"
	}
	return
}
