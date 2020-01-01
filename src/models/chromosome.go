package models

type ChromosomeModel struct {
	Fitness    float64
	Generation uint
	Genes      []CardModel
}

func NewChromosome(size uint) ChromosomeModel {
	return ChromosomeModel{Generation: 1, Genes: make([]CardModel, size)}
}

func (self *ChromosomeModel) GenerateGenes(order OrderModel, stores StoresGroup) {
	count := 0

	for _, orderCard := range order.Cards {
		j := uint(0)

		for j < orderCard.Amount {
			card := stores.FindCardWithStock(orderCard.ID)

			if card != nil {
				self.Genes[count] = card.CardModel
				j++
				count++
				card.DecreaseAmount()
			}
		}
	}

	self.CalculateFitness()
}

func (self *ChromosomeModel) CalculateFitness() {
	self.Fitness = 0

	for _, i := range self.Genes {
		self.Fitness += i.Price
	}
}
