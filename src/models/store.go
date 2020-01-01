package models

type StoreModel struct {
	Name  string           `json:"name"`
	Cards []StoreCardModel `json:"cards"`
}

type StoreCardModel struct {
	Amount uint `json:"amount"`
	CardModel
}
