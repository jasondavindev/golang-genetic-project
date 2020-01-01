package models

type ChromosomeModel struct {
	Fitness    float64
	Generation uint
	Genes      []CardModel
}

func NewChromosome(size uint) ChromosomeModel {
	return ChromosomeModel{Generation: 1, Genes: make([]CardModel, size)}
}

func (self ChromosomeModel) GenerateGenes() {

}
