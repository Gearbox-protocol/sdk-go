package log

import (
	"fmt"
	"log"
	"runtime"
	"testing"

	amqp "github.com/rabbitmq/amqp091-go"
)

var testLogModule *testing.T

func SetTestLogging(t *testing.T) {
	testLogModule = t
}
func Verbosef(msg string, args ...interface{}) {
	log.Printf(DetectFunc()+msg, args...)
}
func Verbose(v ...interface{}) {
	x := []interface{}{DetectFunc()}
	x = append(x, v...)
	log.Println(x...)
}

func Updatef(msg string, args ...interface{}) {
	msgFormat := "[Update] " + DetectFunc() + msg
	amqpSendf(msgFormat, args)
	if testLogModule == nil {
		log.Printf(msgFormat, args...)
	} else {
		testLogModule.Logf(msgFormat, args...)
	}
}

func Update(v ...interface{}) {
	args := []interface{}{"[Update]: " + DetectFunc()}
	args = append(args, v...)
	amqpSend(args)
	if testLogModule == nil {
		log.Println(args...)
	} else {
		testLogModule.Log(args...)
	}
}

func Warnf(msg string, args ...interface{}) {
	msgFormat := "[Warn] " + DetectFunc() + msg
	amqpSendf(msgFormat, args)
	if testLogModule == nil {
		log.Printf(msgFormat, args...)
	} else {
		testLogModule.Logf(msgFormat, args...)
	}
}

func Warn(v ...interface{}) {
	args := []interface{}{"[Warn]: " + DetectFunc()}
	args = append(args, v...)
	amqpSend(args)
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
	amqpSendf(msgFormat, args)
	if testLogModule == nil {
		log.Printf(msgFormat, args...)
	} else {
		testLogModule.Logf(msgFormat, args...)
	}
}

func Error(v ...interface{}) {
	args := []interface{}{"[Error]: " + DetectFunc()}
	args = append(args, v...)
	amqpSend(args)
	if testLogModule == nil {
		log.Println(args...)
	} else {
		testLogModule.Log(args...)
	}
}

func Msgf(msg string, args ...interface{}) {
	amqpSendf(msg, args)
	msgFormat := DetectFunc() + msg
	log.Printf("[AMQP]"+msgFormat, args...)
}

func Msg(v ...interface{}) {
	amqpSend(v)
	args := []interface{}{"[AMQP]" + DetectFunc()}
	args = append(args, v...)
	log.Println(args...)
}

func Fatalf(msg string, args ...interface{}) {
	msgFormat := "[Fatal]: " + DetectFunc() + msg
	amqpSendf(msgFormat, args)
	if testLogModule == nil {
		log.Fatalf(msgFormat, args...)
	} else {
		testLogModule.Fatalf(msgFormat, args...)
	}
}

func Fatal(v ...interface{}) {
	args := []interface{}{"[Fatal]: " + DetectFunc()}
	args = append(args, v...)
	amqpSend(args)
	if testLogModule == nil {
		log.Fatal(args...)
	} else {
		testLogModule.Fatal(args...)
	}
}

func DetectFunc() string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf(" %s:%d ", file, line)
}
func DetectFuncAtStackN(n int) string {
	_, file, line, _ := runtime.Caller(n)
	return fmt.Sprintf(" %s:%d ", file, line)
}

func CheckFatal(err error) {
	args := []interface{}{"[Fatal]: " + DetectFunc(), err}
	if err != nil {
		amqpSend(args)
		if testLogModule == nil {
			log.Fatal(args...)
		} else {
			testLogModule.Fatal(args...)
		}
	}
}

var ch *amqp.Channel
var netName, appName string

func SetAMQP(_ch *amqp.Channel, netName_ string, appName_ string) {
	ch = _ch
	netName = netName_
	appName = appName_
}
func amqpSend(v []interface{}) {
	alert := fmt.Sprint(v...)
	send(alert)
}
func amqpSendf(msg string, args []interface{}) {
	alert := fmt.Sprintf(msg, args...)
	send(alert)
}
func send(message string) {
	if ch == nil {
		return
	}
	err := ch.Publish(
		"TelegramBot", // exchange
		netName,       // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("[%s]%s:", netName, appName) + message),
		})
	if err != nil {
		log.Println("Cant sent notification", err)
	}
}
