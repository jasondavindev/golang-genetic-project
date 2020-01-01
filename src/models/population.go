package models

type PopulationModel struct {
	Size        uint
	Chromosomes []ChromosomeModel
}

func NewPopulation(size uint) *PopulationModel {
	return &PopulationModel{Size: size, Chromosomes: make([]ChromosomeModel, size)}
}

func (self *PopulationModel) GenerateChromossomes(stores []StoreModel) {

}
