package main

import (
	"fmt"

	"github.com/jasondavindev/golang-genetic-project/src/models"
	"github.com/jasondavindev/golang-genetic-project/src/util"
)

func main() {
	var order models.OrderModel
	var stores models.StoresGroup

	LoadOrder("/home/app/files/order.json", &order)
	LoadStores("/home/app/files/stores.json", &stores)

	population := models.NewPopulation(10)
	population.GenerateChromossomes(order, stores)
	population.GenerateChromossomesGenes(order, stores)
	fmt.Println(population.GetTopOne())
}

func LoadOrder(path string, obj *models.OrderModel) interface{} {
	return util.BindFileWith(path, obj)
}

func LoadStores(path string, obj *models.StoresGroup) interface{} {
	return util.BindFileWith(path, obj)
}
