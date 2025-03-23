package log

import (
	"strings"
)

func InitLogging(appType string, chainId int64, cfg CommonEnvs, ethProvider string) {
	routeKey := GetNetworkName(chainId)
	if strings.Contains(ethProvider, "anvil.gearbox.foundation") {
		routeKey = GetTestNet(chainId)
	}
	NewAMQPService(
		cfg.AMQPEnable,
		cfg.AMQPUrl,
		LoggingConfig{
			Exchange:     "TelegramBot",
			ROUTE_KEY:    routeKey,
			RiskEndpoint: cfg.RiskEndpoint,
			RiskSecret:   cfg.RiskSecret,
		},
		cfg.AppName,
	)
	Warn(appType + " started")
}
