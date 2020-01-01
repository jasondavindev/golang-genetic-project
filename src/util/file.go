package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
)

func BindFile(path string, obj interface{}) interface{} {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		panic(err)
	}

	reflectType := reflect.ValueOf(obj).Type()
	model := reflect.New(reflectType).Interface()
	err = json.Unmarshal(byteValue, &model)

	if err != nil {
		panic(err)
	}

	return model
}
