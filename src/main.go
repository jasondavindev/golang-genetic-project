package main

import (
	"fmt"

	"github.com/jasondavindev/golang-genetic-project/src/config"
	"github.com/jasondavindev/golang-genetic-project/src/models"
	"github.com/jasondavindev/golang-genetic-project/src/util"
)

func main() {
	if err := config.InitConfigs(); err != nil {
		panic(err)
	}

	var order models.OrderModel
	var stores models.StoresGroup

	LoadOrder(config.GetConfig().Files.Order, &order)
	LoadStores(config.GetConfig().Files.Stores, &stores)

	population := models.NewPopulation(config.GetConfig().Population.Size)
	population.GenerateChromossomes(order, stores)
	population.GenerateChromossomesGenes(order, stores)

	fmt.Println("Top Fitness before mutation", population.GetTopOne().Fitness)
	population.MakeMutation(&stores, config.GetConfig().Mutation.Cycles)
	fmt.Println("Top Fitness after mutation", population.GetTopOne().Fitness)
}

// LoadOrder load order file
func LoadOrder(path string, obj *models.OrderModel) error {
	return util.BindFileWith(path, obj)
}

// LoadStores load stores file
func LoadStores(path string, obj *models.StoresGroup) error {
	return util.BindFileWith(path, obj)
}
