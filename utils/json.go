package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadFile(fileName string) []byte {
	jsonFile, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	return jsonFile
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

func ReadJsonAndSetInterface(fileName string, data interface{}) {
	reader := bytes.NewReader(ReadFile(fileName))
	ReadJsonReaderAndSetInterface(reader, data)
}

func ReadJsonReaderAndSetInterface(reader io.Reader, data interface{}) {
	d := json.NewDecoder(reader)
	d.UseNumber()
	if err := d.Decode(&data); err != nil {
		fmt.Println("error:", err)
	}
}

func ReadJson(fileName string) map[string]interface{} {
	byteValue := ReadFile(fileName)
	return ReadJsonReader(bytes.NewReader(byteValue))
}
func ReadJsonReader(reader io.Reader) map[string]interface{} {
	data := map[string]interface{}{}
	d := json.NewDecoder(reader)
	d.UseNumber()
	if err := d.Decode(&data); err != nil {
		fmt.Println("error:", err)
	}
	return data
}

func SetJson(byteValue []byte, data interface{}) {
	d := json.NewDecoder(bytes.NewReader(byteValue))
	// use number instead of encoding as float
	d.UseNumber()
	if err := d.Decode(&data); err != nil {
		fmt.Println("error:", err)
	}
}
