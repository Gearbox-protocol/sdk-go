package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
	"testing"
)

var testLogModule *testing.T

func SetTestLogging(t *testing.T) {
	testLogModule = t
}
func severityFormat(severity string) string {
	return "[" + severity + "]: "
}
func printf(severity LEVEL, msg string, args ...interface{}) string {
	if severity < logLevel {
		return ""
	}
	_log := fmt.Sprintf(severityFormat(severity.ToString())+DetectFuncAtStackN(3)+msg, args...)
	if testLogModule == nil {
		log.Println(_log)
	} else {
		testLogModule.Log(_log)
	}
	return fmt.Sprintf(severityFormat(severity.ToString())+msg, args...)
}
func println(severity LEVEL, args ...interface{}) string {
	if severity < logLevel {
		return ""
	}
	return printlnStr(severity.ToString(), 4, args...)
}
func printlnStr(severity string, depth int, args ...interface{}) string {
	_log := severityFormat(severity) + DetectFuncAtStackN(depth) + fmt.Sprintln(args...)
	if testLogModule == nil {
		log.Printf(_log)
	} else {
		testLogModule.Logf(_log)
	}
	return severityFormat(severity) + fmt.Sprintln(args...)
}

// ---- Don't send to amqp
// trace
// debug
// info
// ----- To channel Alerts
// warn
// error
// msg
// fatal
// ----- Risk condition alerts
// Alert
func Tracef(msg string, args ...interface{}) {
	printf(TRACE, msg, args...)
}
func Trace(args ...interface{}) {
	println(TRACE, args...)
}
func TraceAtN(depth int, args ...interface{}) {
	if TRACE < logLevel {
		return
	}
	printlnStr(TRACE.ToString(), depth, args...)
}
func Debugf(msg string, args ...interface{}) {
	printf(DEBUG, msg, args...)
}
func Debug(args ...interface{}) {
	println(DEBUG, args...)
}

func Warnf(msg string, args ...interface{}) {
	if _log := printf(WARN, msg, args...); _log != "" {
		send(_log, WARN)
	}
}

func Warn(args ...interface{}) {
	if _log := println(WARN, args...); _log != "" {
		send(_log, WARN)
	}
}

func Infof(msg string, args ...interface{}) {
	printf(INFO, msg, args...)
}

func Info(args ...interface{}) {
	println(INFO, args...)
}

func InfoStackN(n int, v ...interface{}) {
	args := []interface{}{"[Info]: " + DetectFuncAtStackN(n)}
	args = append(args, v...)
	if testLogModule == nil {
		log.Println(args...)
	} else {
		testLogModule.Log(args...)
	}
}

func Errorf(msg string, args ...interface{}) {
	if _log := printf(ERROR, msg, args...); _log != "" {
		send(_log, ERROR)
	}
}

func Error(args ...interface{}) {
	if _log := println(ERROR, args...); _log != "" {
		send(_log, ERROR)
	}
}

func Fatalf(msg string, args ...interface{}) {
	debug.PrintStack()
	if _log := printf(FATAL, msg, args...); _log != "" {
		send(_log, FATAL)
	}
	os.Exit(1)
}

func Fatal(args ...interface{}) {
	debug.PrintStack()
	if _log := println(FATAL, args...); _log != "" {
		send(_log, FATAL)
	}
	os.Exit(1)
}

func CheckFatal(err error) {
	if err != nil {
		debug.PrintStack()
		msg := "[Fatal]: " + DetectFuncAtStackN(2) + err.Error()
		if testLogModule == nil {
			log.Println(msg)
		} else {
			testLogModule.Log(msg)
		}
		send(msg, FATAL)
		os.Exit(1)
	}
}

func AMQPMsgf(msg string, args ...interface{}) {
	_log := printf(AMQP, msg, args...)
	send(_log, AMQP)
}

func AMQPMsg(args ...interface{}) {
	_log := println(AMQP, args...)
	send(_log, AMQP)
}

// //
var cwdLen int

func DetectFuncAtStackN(n int) string {
	if cwdLen == 0 {
		s, _ := os.Getwd()
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == '/' {
				cwdLen = i + 1
				break
			}
		}
	}
	_, file, line, _ := runtime.Caller(n)
	if ind := strings.IndexRune(file, '@'); ind == -1 {
		return fmt.Sprintf(" %s:%d ", file[cwdLen:], line)
	} else {
		remainingPath := file[ind+1:]
		extraInd := strings.IndexRune(remainingPath, '/')
		return fmt.Sprintf(" %s:%d ", remainingPath[extraInd+1:], line)
	}
}

func WrapErrWithLine(err error) error {
	return fmt.Errorf("%s: %v", DetectFuncAtStackN(2), err)
}
func WrapErrWithLineN(err error, n int) error {
	return fmt.Errorf("%s: %v", DetectFuncAtStackN(n), err)
}
