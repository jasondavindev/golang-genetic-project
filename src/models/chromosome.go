package models

import (
	"github.com/jasondavindev/golang-genetic-project/src/util"
)

// ChromosomeModel model of chromossomes
type ChromosomeModel struct {
	Fitness    float64
	Generation uint
	Genes      []CardModel
}

// NewChromosome creates new chromossome
func NewChromosome(size uint) ChromosomeModel {
	return ChromosomeModel{Generation: 1, Genes: make([]CardModel, size)}
}

// NewGeneration creates new chromossome based on two parent chromossomes
func NewGeneration(parents []ChromosomeModel) ChromosomeModel {
	children := NewChromosome(uint(len(parents[0].Genes)))

	for k := range children.Genes {
		if parents[0].Genes[k].Price > parents[1].Genes[k].Price {
			children.Genes[k] = parents[1].Genes[k]
		} else {
			children.Genes[k] = parents[0].Genes[k]
		}
	}

	return children
}

// GenerateGenes Generates genes of chromossome randomly
func (chromo *ChromosomeModel) GenerateGenes(order OrderModel, stores StoresGroup) {
	count := 0

	for _, orderCard := range order.Cards {
		j := uint(0)

		for j < orderCard.Amount {
			card := stores.FindCardWithStock(orderCard.ID)

			if card != nil {
				chromo.Genes[count] = card.CardModel
				j++
				count++
				card.DecreaseAmount()
			}
		}
	}

	chromo.CalculateFitness()
}

// CalculateFitness sum the price of all chromosome cards
func (chromo *ChromosomeModel) CalculateFitness() {
	chromo.Fitness = 0

	for _, i := range chromo.Genes {
		chromo.Fitness += i.Price
	}
}

// Mutation Mutates chromosome genes. N genes (cards)
// are chosen randomly and then replaced by other lower priced random cards
func (chromo *ChromosomeModel) Mutation(chance int, genes int, stores *StoresGroup) {
	random := util.Rand(100)

	if random > chance {
		return
	}

	for i := 1; i <= genes; i++ {
		rand1 := util.Rand(len(chromo.Genes))
		randomGene := chromo.Genes[rand1]

		rand2 := util.Rand(len(*stores))
		randomStore := (*stores)[rand2]

		var card *StoreCardModel

		for card = randomStore.FindCardWithStock(randomGene.ID); card == nil; card = randomStore.FindCardWithStock(randomGene.ID) {
			rand2 = util.Rand(len(*stores))
			randomStore = (*stores)[rand2]
		}

		if card.Price < randomGene.Price {
			chromo.Genes[rand1] = card.CardModel
			card.DecreaseAmount()
		}
	}
}
