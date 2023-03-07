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
func printf(severity, msg string, args ...interface{}) string {
	_log := fmt.Sprintf(severityFormat(severity)+DetectFuncAtStackN(3)+msg, args...)
	if testLogModule == nil {
		log.Println(_log)
	} else {
		testLogModule.Log(_log)
	}
	return fmt.Sprintf(severityFormat(severity)+msg, args...)
}
func println(severity string, args ...interface{}) string {
	_log := severityFormat(severity) + DetectFuncAtStackN(3) + fmt.Sprintln(args...)
	if testLogModule == nil {
		log.Printf(_log)
	} else {
		testLogModule.Logf(_log)
	}
	return severityFormat(severity) + fmt.Sprintln(args...)
}

// ---- Don't send to amqp
// verbose
// info
// ----- To channel Alerts
// warn
// error
// msg
// fatal
// ----- Risk condition alerts
// Alert
func Verbosef(msg string, args ...interface{}) {
	printf("Debug", msg, args...)
}
func Verbose(args ...interface{}) {
	println("Debug", args...)
}

//
func Warnf(msg string, args ...interface{}) {
	_log := printf("Warn", msg, args...)
	send(false, _log)
}

func Warn(args ...interface{}) {
	_log := println("Warn", args...)
	send(false, _log)
}

func Infof(msg string, args ...interface{}) {
	printf("Info", msg, args...)
}

func Info(args ...interface{}) {
	println("Info", args...)
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
	_log := printf("Error", msg, args...)
	send(false, _log)
}

func Error(args ...interface{}) {
	_log := println("Error", args...)
	send(false, _log)
}

func Fatalf(msg string, args ...interface{}) {
	debug.PrintStack()
	_log := printf("Fatal", msg, args...)
	send(false, _log)
	os.Exit(1)
}

func Fatal(args ...interface{}) {
	debug.PrintStack()
	_log := println("Fatal", args...)
	send(false, _log)
	os.Exit(1)
}

func CheckFatal(err error) {
	if err != nil {
		debug.PrintStack()
		msg := "[Fatal]: " + DetectFuncAtStackN(2) + err.Error()
		if testLogModule == nil {
			log.Fatalln(msg)
		} else {
			testLogModule.Fatal(msg)
		}
		send(false, msg)
	}
}

func AMQPMsgf(msg string, args ...interface{}) {
	_log := printf("AMQP", msg, args...)
	send(false, _log)
}

func AMQPMsg(args ...interface{}) {
	_log := println("AMQP", args...)
	send(false, _log)
}

////
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
