package log

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
)

var _logConfig LoggingConfig

type LoggingConfig struct {
	// for amqp
	Exchange string `json:"-"`
	// for risk service
	RiskEndpoint string `json:"-"`
	RiskSecret   string `json:"-"`
	//
	// for chain for risk msg
	ChainId int64 `json:"chain"`
	appDetails
}

func (cfg LoggingConfig) getRiskToken() string {
	token := crypto.Keccak256Hash([]byte(cfg.App + cfg.Instance + cfg.RiskSecret))
	return token.String()[2:]
}

type CommonEnvs struct {
	//
	AppName string `env:"APP_NAME" default:"Third-eye"`
	// amqp url and enable flags
	AMQPUrl    string `env:"CLOUDAMQP_URL" validate:"required"`
	AMQPEnable string `env:"AMQP_ENABLE" validate:"required"`
	// for risk service
	RiskEndpoint string `env:"RISK_ENDPOINT"`
	RiskSecret   string `env:"RISK_SECRET"`
}
type appDetails struct {
	// id of application monitoring
	App string `json:"monitor_code"`
	// operator of the application, H or K
	Instance string `json:"instance"`
}

type RiskHeader struct {
	BlockNumber int64 `json:"block_number,omitempty"`
	// condition to monitor
	EventCode string `json:"event_code"`
}
type RiskAlert struct {
	RiskHeader
	FirstOccuredId string `json:"first_occured,omitempty"`
	Id             string `json:"id"`
	TxHash         string `json:"tx_hash,omitempty"`
	// message to send
	Msg string `json:"message"`
}

type riskMsgFullScheme struct {
	RiskAlert
	LoggingConfig
}

func SendRiskAlert(event RiskAlert) {
	_log := printlnStr(event.EventCode, 3, event.Msg)
	if event.Id == "" {
		event.Id = uuid.New().String()
	}
	postReqToRisk(riskMsgFullScheme{
		RiskAlert:     event,
		LoggingConfig: _logConfig,
	})
	send(true, _log)
}

func postReqToRisk(alert riskMsgFullScheme) {
	if _logConfig.RiskEndpoint == "" {
		return
	}
	body, err := json.Marshal([]riskMsgFullScheme{alert})
	if err != nil {
		Error(err)
		return
	}
	req, _ := http.NewRequest(http.MethodPost, _logConfig.RiskEndpoint, bytes.NewBuffer(body))
	if _logConfig.RiskSecret != "" {
		req.Header.Set("Authorization", "Bearer "+_logConfig.getRiskToken())
		req.Header.Set("content-type", "application/json")
	}
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		Error(err)
	} else if resp.StatusCode != 200 {
		var body string
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			body = err.Error()
		} else {
			body = string(data)
		}
		Error("Risk service returns error: ", resp.StatusCode, body)
	}
}

type RiskMsgTimer struct {
	FirstOccuredId string
	lastTs         time.Time
}

type TimerFn func(time.Duration) (timer *RiskMsgTimer, canSendMsg bool, isFirstMsg bool)

func GetRiskMsgTimer() TimerFn {
	timer := &RiskMsgTimer{}
	return func(duration time.Duration) (*RiskMsgTimer, bool, bool) {
		if timer.lastTs.IsZero() {
			timer.FirstOccuredId = uuid.New().String()
			timer.lastTs = time.Now()
			return timer, true, true
		}
		if time.Since(timer.lastTs) > duration {
			timer.lastTs = time.Now()
			return timer, true, false
		}
		return timer, false, false
	}
}

func SendRiskAlertPerTimer(alert RiskAlert, timerFn TimerFn, interval time.Duration) {
	if interval == 0 {
		SendRiskAlert(alert)
		return
	}
	if timer, canSendMsg, isFirstMsg := timerFn(interval); canSendMsg {
		alert.FirstOccuredId = timer.FirstOccuredId
		if isFirstMsg {
			alert.Id = timer.FirstOccuredId
		}
		SendRiskAlert(alert)
	}
}
