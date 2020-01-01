package models

type OrderModel struct {
	Cards []StoreCardModel `json:"cards"`
}

func (self *OrderModel) GeneratePopulation(size uint) *PopulationModel {
	return NewPopulation(size)
}
