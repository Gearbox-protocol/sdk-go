package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(fileName string) []byte {
	jsonFile, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	return jsonFile
}

func ReadJsonAndSet(fileName string) []map[string]interface{} {
	data := []map[string]interface{}{}
	byteValue := ReadFile(fileName)
	d := json.NewDecoder(bytes.NewReader(byteValue))
	d.UseNumber()
	if err := d.Decode(&data); err != nil {
		fmt.Println("error:", err)
	}
	return data
}
func ReadJsonAndSetInterface(fileName string, data interface{}) {
	byteValue := ReadFile(fileName)
	d := json.NewDecoder(bytes.NewReader(byteValue))
	d.UseNumber()
	if err := d.Decode(&data); err != nil {
		fmt.Println("error:", err)
	}
}

func ReadJson(fileName string) map[string]interface{} {
	data := map[string]interface{}{}
	byteValue := ReadFile(fileName)
	d := json.NewDecoder(bytes.NewReader(byteValue))
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
