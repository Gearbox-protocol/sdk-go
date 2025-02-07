package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"

	"github.com/Gearbox-protocol/sdk-go/log"
)

func ReadFile(fileName string) ([]byte, error) {
	jsonFile, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return jsonFile, nil
}

// func ReadJsonAndSet(fileName string) []map[string]interface{} {
// 	data := []map[string]interface{}{}
// 	byteValue := ReadFile(fileName)
// 	d := json.NewDecoder(bytes.NewReader(byteValue))
// 	d.UseNumber()
// 	if err := d.Decode(&data); err != nil {
// 		fmt.Println("error:", err)
// 	}
// 	return data
// }

func ReadJsonAndSetInterface(fileName string, data interface{}) error {
	_bytes, err := ReadFile(fileName)
	if err != nil {
		log.Debug(log.WrapErrWithLineN(err, 3))
		return err
	}
	reader := bytes.NewReader(_bytes)
	return ReadJsonReaderAndSetInterface(reader, data)
}

func ReadJsonReaderAndSetInterface(reader io.Reader, data interface{}) error {
	d := json.NewDecoder(reader)
	d.UseNumber()
	if err := d.Decode(&data); err != nil {
		log.Debug("error:", log.WrapErrWithLineN(err, 3))
		return err
	}
	return nil
}

// map passed with reference can't be expected to have data filled in it
// that's why we are getting the data as return variable
func ReadJson(fileName string) (map[string]interface{}, error) {
	byteValue, err := ReadFile(fileName)
	if err != nil {
		return nil, log.WrapErrWithLine(err)
	}
	return ReadJsonReader(bytes.NewReader(byteValue))
}
func ReadJsonReader(reader io.Reader) (map[string]interface{}, error) {
	data := map[string]interface{}{}
	d := json.NewDecoder(reader)
	d.UseNumber()
	if err := d.Decode(&data); err != nil {
		fmt.Println("error:", err)
		return nil, err
	}
	return data, nil
}

func SetJson(byteValue []byte, data interface{}) error {
	d := json.NewDecoder(bytes.NewReader(byteValue))
	// use number instead of encoding as float
	d.UseNumber()
	if err := d.Decode(&data); err != nil {
		fmt.Println("error:", log.WrapErrWithLineN(err, 2))
		return err
	}
	return nil
}

func InterfaceToInt(v interface{}) (int64, error) {
	switch ans := v.(type) {
	case string:
		num, err := strconv.ParseInt(ans, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("err(%v) while converting %s to int", err, ans)
		}
		return num, nil
	case int64, int, int32, int16:
		return reflect.ValueOf(ans).Int(), nil
	case float64:
		return int64(ans), nil
	case json.Number:
		num, err := ans.Int64()
		if err != nil {
			f, err := ans.Float64()
			return int64(f), err
		}
		return num, err
	default:
		return 0, fmt.Errorf("can't convert %v to int", v)
	}
}
