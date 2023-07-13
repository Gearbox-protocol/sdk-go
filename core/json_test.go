package core

import (
	"encoding/json"
	"testing"
)

func TestJsonFloatMapUnmarshal(t *testing.T) {
	str := []byte(`{"1":1}`)
	msg := JsonFloatMap{}
	if err := json.Unmarshal(str, &msg); err != nil {
		t.Fatal(err)
	}
	if msg["1"] != 1 {
		t.Fatal("unmarshal failed")
	}
}

func TestJsonFloatMapCopy(t *testing.T) {
	msg := JsonFloatMap{"1": 1}
	bytes, err := json.Marshal(msg)
	if err != nil {
		t.Fatal(err)
	}
	msg2 := msg.Copy()
	bytes2, err := json.Marshal(msg2)
	if err != nil {
		t.Fatal(err)
	}
	if string(bytes) != string(bytes2) {
		t.Fatal("copy failed")
	}
}

func TestJsonFloatMapValueInUSD(t *testing.T) {
	msg := JsonFloatMap{"1": 1}
	if msg.ValueInUSD(map[string]float64{"1": 10}) != 10 {
		t.Fatal("value in USD failed")
	}
}

func TestJsonFloatMapScan(t *testing.T) {
	msg := JsonFloatMap{}
	msgStr := `{"1":1}`
	if err := msg.Scan(msgStr); err != nil {
		t.Fatal(err)
	}
	if msg["1"] != 1 {
		t.Fatal("unmarshal failed")
	}
	msg = JsonFloatMap{}
	if err := msg.Scan([]byte(msgStr)); err != nil {
		t.Fatal(err)
	}
	if msg["1"] != 1 {
		t.Fatal("unmarshal failed")
	}
}

// //////
func TestJsonUnmarshal(t *testing.T) {
	msg := Json{}
	err := json.Unmarshal([]byte(`{"1":1}`), &msg)
	if err != nil {
		t.Fatal(err)
	}
}
