package log

import (
	"testing"
	"time"
)

func TestTimerFn(t *testing.T) {
	timerFn := GetRiskMsgTimer()

	dur := 1 * time.Second
	// test if the id is set on first call and can send msg
	timer, canSendMsg, isFirstMsg := timerFn(dur)
	firstId := timer.FirstOccuredId
	if canSendMsg != true || isFirstMsg != true || firstId == "" {
		t.Error("timerFn() should return {!=''}, true, true for the first time")
	}
	// test if within the interval from lastSendMsg can't send msg
	time.Sleep(time.Second / 2)
	timer, canSendMsg, isFirstMsg = timerFn(dur)
	if canSendMsg != false || isFirstMsg != false || timer.FirstOccuredId != firstId {
		t.Error("timerFn() should return {firstId}, false, false within the interval from first call")
	}
	// test if out of the interval from lastSendMsg can send msg
	time.Sleep(time.Second)
	timer, canSendMsg, isFirstMsg = timerFn(dur)
	if canSendMsg != true || isFirstMsg != false || timer.FirstOccuredId != firstId {
		t.Error("timerFn() should return {firstId}, true, false out of the interval from first call")
	}
	// test if within the interval from lastSendMsg can't send msg
	time.Sleep(time.Second / 2)
	timer, canSendMsg, isFirstMsg = timerFn(dur)
	if canSendMsg != false || isFirstMsg != false || timer.FirstOccuredId != firstId {
		t.Error("timerFn() should return {firstId}, false, false within the interval from lastSendMsg")
	}
}
