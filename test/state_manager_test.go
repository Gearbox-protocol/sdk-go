package test

import (
	"testing"

	"github.com/Gearbox-protocol/sdk-go/log"
)

func TestBinarySearch(t *testing.T) {
	log.SetTestLogging(t)
	obj := []*NumAndData{
		{1, "a"},
		{5, "d"},
	}
	lastEntryTillBlock(1, obj)

	if lastEntryTillBlock(1, obj).BlockNum != 1 {
		t.Fatal("failed")
	}
	if lastEntryTillBlock(0, obj) != nil {
		t.Fatal("failed")
	}
	if lastEntryTillBlock(6, obj).BlockNum != 5 {
		t.Fatal("failed")
	}
}
