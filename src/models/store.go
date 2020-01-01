package models

import (
	"math/rand"
	"time"
)

type StoreModel struct {
	Name  string           `json:"name"`
	Cards []StoreCardModel `json:"cards"`
}

type StoreCardModel struct {
	Amount uint `json:"amount"`
	CardModel
}

type StoresGroup []StoreModel

func (self *StoreModel) FindCardWithStock(cardID uint) *StoreCardModel {
	for _, card := range self.Cards {
		if card.Amount > 0 && card.ID == cardID {
			return &card
		}
	}

	return nil
}

func (self *StoresGroup) FindCardWithStock(cardID uint) *StoreCardModel {
	s1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	var card *StoreCardModel

	for card == nil {
		random := s1.Intn(len(*self))
		card = (*self)[random].FindCardWithStock(cardID)
	}

	return card
}

func (self *StoreCardModel) DecreaseAmount() {
	self.Amount--
}
