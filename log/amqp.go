package log

import (
	"log"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
)

var _amqpChannel *amqp.Channel

// Service constructor
func NewAMQPService(amqpEnable, amqpUrl string, logConfig LoggingConfig, appName string) {
	_logConfig = logConfig
	splits := strings.SplitN(appName, "@", 2)
	_logConfig.App = splits[0]
	if len(splits) > 1 {
		_logConfig.Instance = splits[1]
	}

	if amqpEnable == "0" {
		return
	}
	//
	conn, err := amqp.Dial(amqpUrl)
	if err != nil {
		Error(err, "Failed to connect to RabbitMQ")
	}
	ch, err := conn.Channel()
	if err != nil {
		Error(err, "Failed to open a channel")
	}
	//
	_amqpChannel = ch
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

func send(important bool, message string) {
	if _amqpChannel == nil {
		return
	}
	err := _amqpChannel.Publish(
		_logConfig.Exchange,                // exchange
		GetNetworkName(_logConfig.ChainId), // routing key
		false,                              // mandatory
		false,                              // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
			Headers:     amqp.Table{"important": important},
			AppId:       _logConfig.App,
			// UserId:      _logConfig.Instance,
		})
	if err != nil {
		log.Println("Cant sent notification", err)
	}
}
