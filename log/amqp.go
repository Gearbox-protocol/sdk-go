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
func NewAMQPService(chainId int64, AMPQEnable string, AMQPURL, appName string) {
	if AMPQEnable == "0" {
		return
	}
	conn, err := amqp.Dial(AMQPURL)
	failOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	SetAMQP(ch, GetNetworkName(chainId), appName)
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
