package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
)

// BindFile returns a structed file
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

// BindFileWith structures a file and bind it to obj
func BindFileWith(path string, obj interface{}) error {
	file, err := os.Open(path)

	if err != nil {
		return err
	}

	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, obj)

	if err != nil {
		return err
	}

	return nil
}
