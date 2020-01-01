package models

import (
	"github.com/jasondavindev/golang-genetic-project/src/util"
)

// StoreModel model of store
type StoreModel struct {
	Name  string           `json:"name"`
	Cards []StoreCardModel `json:"cards"`
}

// StoreCardModel model of cards in store file
type StoreCardModel struct {
	Amount uint `json:"amount"`
	CardModel
}

// StoresGroup slice of stores
type StoresGroup []StoreModel

// FindCardWithStock Finds passed card with stock in a store
func (store *StoreModel) FindCardWithStock(cardID uint) *StoreCardModel {
	for _, card := range store.Cards {
		if card.Amount > 0 && card.ID == cardID {
			return &card
		}
	}

	return nil
}

// FindCardWithStock Finds passed card with stock in all stores
func (stores *StoresGroup) FindCardWithStock(cardID uint) *StoreCardModel {
	var card *StoreCardModel

	for card == nil {
		random := util.Rand(len(*stores))
		card = (*stores)[random].FindCardWithStock(cardID)
	}

	return card
}

// DecreaseAmount decreases the amount in stock of a card
func (card *StoreCardModel) DecreaseAmount() {
	card.Amount--
}
