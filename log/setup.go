package log

import (
	"strings"
)

var _routeKey NETWORK

func TestNetOrMainnet() NETWORK {
	return _routeKey
}
func InitLogging(appType string, basechainId int64, cfg CommonEnvs, ethProvider string) {
	_routeKey = GetNetworkName(basechainId)
	if strings.Contains(ethProvider, "anvil.gearbox.foundation") {
		_routeKey = GetTestNet(basechainId)
	}
	NewAMQPService(
		cfg.AMQPEnable,
		cfg.AMQPUrl,
		LoggingConfig{
			Exchange:     "TelegramBot",
			ROUTE_KEY:    _routeKey,
			RiskEndpoint: cfg.RiskEndpoint,
			RiskSecret:   cfg.RiskSecret,
		},
		cfg.AppName,
	)
	Warn(appType + " started")
}
