package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

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
		fmt.Println(log.WrapErrWithLineN(err, 2))
		return err
	}
	reader := bytes.NewReader(_bytes)
	return ReadJsonReaderAndSetInterface(reader, data)
}

func ReadJsonReaderAndSetInterface(reader io.Reader, data interface{}) error {
	d := json.NewDecoder(reader)
	d.UseNumber()
	if err := d.Decode(&data); err != nil {
		fmt.Println("error:", log.WrapErrWithLineN(err, 2))
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
