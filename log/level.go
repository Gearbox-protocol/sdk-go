package log

import (
	"os"
)

var logLevel LEVEL

type LEVEL int

const (
	TRACE LEVEL = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	AMQP
)

func toString(lvl LEVEL) string {
	switch lvl {
	case TRACE:
		return "TRACE"
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	case AMQP:
		return "AMQP"
	}
	return ""
}

func init() {
	lvl, ok := os.LookupEnv("GB_LOG_LEVEL")
	if !ok {
		lvl = "INFO"
	}
	switch lvl {
	case "TRACE":
		logLevel = TRACE
	case "DEBUG":
		logLevel = DEBUG
	case "INFO":
		logLevel = INFO
	case "WARN":
		logLevel = WARN
	case "ERROR":
		logLevel = ERROR
	case "FATAL":
		logLevel = FATAL
	}
}
