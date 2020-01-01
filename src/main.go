package main

import (
	"fmt"

	"github.com/jasondavindev/golang-genetic-project/src/models"
	"github.com/jasondavindev/golang-genetic-project/src/util"
)

func main() {
	order := LoadOrder("/home/app/files/order.json")
	stores := LoadStores("/home/app/files/stores.json")
	population := models.NewPopulation(100)
	fmt.Printf("order %v\nType: %T\n", order, order)
	fmt.Printf("stores %v\nType: %T\n", stores, stores)
	fmt.Println(population)
}

func LoadOrder(path string) interface{} {
	return util.BindFile(path, models.OrderModel{})
}

func LoadStores(path string) interface{} {
	return util.BindFile(path, []models.StoreModel{})
}
