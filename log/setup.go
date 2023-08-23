package log

func InitLogging(appType string, chainId int64, cfg CommonEnvs) {
	NewAMQPService(
		cfg.AMQPEnable,
		cfg.AMQPUrl,
		LoggingConfig{
			Exchange:     "TelegramBot",
			ChainId:      chainId,
			RiskEndpoint: cfg.RiskEndpoint,
			RiskSecret:   cfg.RiskSecret,
		},
		cfg.AppName,
	)
	AMQPMsg(appType + " started")
}
