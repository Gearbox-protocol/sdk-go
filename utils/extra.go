package utils

import "github.com/Gearbox-protocol/sdk-go/log"

func ConvertToListOfInt64(list interface{}) (parsedInts []int64) {
	switch ints := list.(type) {
	case []interface{}:
		for _, _int := range ints {
			var parsedInt int64
			switch parsedV := _int.(type) {
			case int64:
				parsedInt = parsedV
			case float64:
				parsedInt = int64(parsedV)
			default:
				log.Fatalf("QueryPriceFeed token start/end block_num not in int format %v", _int)
			}
			parsedInts = append(parsedInts, parsedInt)
		}
	case []int64:
		parsedInts = ints
	}
	return
}

func ListOfInt64List(_outList interface{}) (ans [][]int64) {
	switch outList := _outList.(type) {
	case []interface{}:
		for _, _inList := range outList {
			ans = append(ans, ConvertToListOfInt64(_inList))
		}
	case [][]int64:
		ans = outList
	}
	return
}
