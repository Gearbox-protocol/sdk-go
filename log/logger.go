package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/debug"
	"strings"
	"testing"

	amqp "github.com/rabbitmq/amqp091-go"
)

var testLogModule *testing.T

func SetTestLogging(t *testing.T) {
	testLogModule = t
}

// verbose
// info
// ---- Logs
// warn
// error
// msg
// ----- Alerts
// fatal
func Verbosef(msg string, args ...interface{}) {
	if testLogModule == nil {
		log.Printf(DetectFunc()+msg, args...)
	} else {
		testLogModule.Logf(DetectFunc()+msg, args...)
	}
}
func Verbose(v ...interface{}) {
	x := []interface{}{DetectFunc()}
	x = append(x, v...)
	if testLogModule == nil {
		log.Println(x...)
	} else {
		testLogModule.Log(x...)
	}
}

func Warnf(msg string, args ...interface{}) {
	msgFormat := "[Warn] " + DetectFunc() + msg
	amqpSendf(true, msgFormat, args)
	if testLogModule == nil {
		log.Printf(msgFormat, args...)
	} else {
		testLogModule.Logf(msgFormat, args...)
	}
}

func Warn(v ...interface{}) {
	args := []interface{}{"[Warn]: " + DetectFunc()}
	args = append(args, v...)
	amqpSend(true, args)
	if testLogModule == nil {
		log.Println(args...)
	} else {
		testLogModule.Log(args...)
	}
}

func Infof(msg string, args ...interface{}) {
	msg = "[Info]: " + DetectFunc() + msg
	if testLogModule == nil {
		log.Printf(msg, args...)
	} else {
		testLogModule.Logf(msg, args...)
	}

}

func Info(v ...interface{}) {
	args := []interface{}{"[Info]: " + DetectFunc()}
	args = append(args, v...)
	if testLogModule == nil {
		log.Println(args...)
	} else {
		testLogModule.Log(args...)
	}
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
	msgFormat := "[Error]: " + DetectFunc() + msg
	amqpSendf(true, msgFormat, args)
	if testLogModule == nil {
		log.Printf(msgFormat, args...)
	} else {
		testLogModule.Logf(msgFormat, args...)
	}
}

func Error(v ...interface{}) {
	args := []interface{}{"[Error]: " + DetectFunc()}
	args = append(args, v...)
	amqpSend(true, args)
	if testLogModule == nil {
		log.Println(args...)
	} else {
		testLogModule.Log(args...)
	}
}

func Msgf(msg string, args ...interface{}) {
	amqpSendf(true, msg, args)
	msgFormat := DetectFunc() + msg
	log.Printf("[AMQP]"+msgFormat, args...)
}

func Msg(v ...interface{}) {
	amqpSend(true, v)
	args := []interface{}{"[AMQP]" + DetectFunc()}
	args = append(args, v...)
	log.Println(args...)
}

func Fatalf(msg string, args ...interface{}) {
	debug.PrintStack()
	msgFormat := "[Fatal]: " + DetectFunc() + msg
	amqpSendf(false, msgFormat, args)
	if testLogModule == nil {
		log.Fatalf(msgFormat, args...)
	} else {
		testLogModule.Fatalf(msgFormat, args...)
	}
}

func Fatal(v ...interface{}) {
	debug.PrintStack()
	args := []interface{}{"[Fatal]: " + DetectFunc()}
	args = append(args, v...)
	amqpSend(false, args)
	if testLogModule == nil {
		log.Fatalln(args...)
	} else {
		testLogModule.Fatal(args...)
	}
}

func CheckFatal(err error) {
	args := []interface{}{"[Fatal]: " + DetectFunc(), err}
	if err != nil {
		amqpSend(false, args)
		if testLogModule == nil {
			log.Fatalln(args...)
		} else {
			testLogModule.Fatal(args...)
		}
	}
}

var _logConfig LoggingConfig

type LoggingConfig struct {
	Network  string
	Exchange string
	App      string
	Channel  *amqp.Channel
}

func amqpSend(important bool, v []interface{}) {
	alert := fmt.Sprint(v...)
	send(important, alert)
}
func amqpSendf(important bool, msg string, args []interface{}) {
	alert := fmt.Sprintf(msg, args...)
	send(important, alert)
}
func send(important bool, message string) {
	if _logConfig.Channel == nil {
		return
	}
	err := _logConfig.Channel.Publish(
		_logConfig.Exchange, // exchange
		_logConfig.Network,  // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(_logConfig.App + ": " + message),
			Headers:     amqp.Table{"important": important},
		})
	if err != nil {
		log.Println("Cant sent notification", err)
	}
}

////
var cwdLen int

func DetectFunc() string {
	if cwdLen == 0 {
		s, _ := os.Getwd()
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == '/' {
				cwdLen = i + 1
				break
			}
		}

	}
	_, file, line, _ := runtime.Caller(2)
	if ind := strings.IndexRune(file, '@'); ind == -1 {
		return fmt.Sprintf(" %s:%d ", file[cwdLen:], line)
	} else {
		remainingPath := file[ind+1:]
		extraInd := strings.IndexRune(remainingPath, '/')
		return fmt.Sprintf(" %s:%d ", remainingPath[extraInd+1:], line)
	}
}
func DetectFuncAtStackN(n int) string {
	_, file, line, _ := runtime.Caller(n)
	return fmt.Sprintf(" %s:%d ", file, line)
}
