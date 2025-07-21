package log

import (
	"fmt"
	"log"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var _amqpChannel *amqp.Channel
var _amqpUrl string

func Elapsed(what string) func() {
	start := time.Now()
	return func() {
		InfoStackN(3, fmt.Sprintf("%s took %v", what, time.Since(start)))
	}
}

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
	_amqpUrl = amqpUrl
	setChannel()
}

func setChannel() {
	conn, err := amqp.Dial(_amqpUrl)
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

func _contains(s []NETWORK, e NETWORK) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func send(message string, alertType LEVEL, important ...bool) {
	if _amqpChannel == nil {
		return
	}

	if _logConfig.ROUTE_KEY == ANVIL {
		_logConfig.ROUTE_KEY = "GOERLI"
	}
	if _contains([]NETWORK{ETHERLINK, LISK, HEMIBTC}, _logConfig.ROUTE_KEY) {
		_logConfig.ROUTE_KEY = "PROD"
	}
	if _contains([]NETWORK{ETHERLINKTEST, LISKTEST, HEMIBTCTEST}, _logConfig.ROUTE_KEY) {
		_logConfig.ROUTE_KEY = "TEST"
	}
	for i := 0; i < 2; i++ {
		err := _amqpChannel.Publish(
			_logConfig.Exchange,          // exchange
			string(_logConfig.ROUTE_KEY), // routing key
			false,                        // mandatory
			false,                        // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
				Headers: amqp.Table{"important": func() bool {
					if len(important) == 0 {
						return false
					}
					return important[0]
				}(), "type": int(alertType)},
				AppId: _logConfig.App,
				// UserId:      _logConfig.Instance,
			})
		if err != nil {
			log.Println("Cant sent notification", err)
			setChannel()
		} else {
			return
		}
	}
}
