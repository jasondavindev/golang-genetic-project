package models

import "math"

type PopulationModel struct {
	Size        uint
	Chromosomes []ChromosomeModel
}

func NewPopulation(size uint) *PopulationModel {
	return &PopulationModel{Size: size, Chromosomes: make([]ChromosomeModel, size)}
}

func (self *PopulationModel) GenerateChromossomes(order OrderModel, stores StoresGroup) {
	for k := range self.Chromosomes {
		self.Chromosomes[k] = NewChromosome(order.FinalAmount)
	}
}

func (self *PopulationModel) GenerateChromossomesGenes(order OrderModel, stores StoresGroup) {
	for idx := range self.Chromosomes {
		self.Chromosomes[idx].GenerateGenes(order, stores)
	}
}

func (self *PopulationModel) GetTopOne() ChromosomeModel {
	top := ChromosomeModel{Fitness: math.MaxFloat64}

	for idx := range self.Chromosomes {
		if self.Chromosomes[idx].Fitness < top.Fitness {
			top = self.Chromosomes[idx]
		}
	}

	return top
}
