package models

import (
	"math"

	"github.com/jasondavindev/golang-genetic-project/src/config"
	"github.com/jasondavindev/golang-genetic-project/src/util"
)

// PopulationModel model of population
type PopulationModel struct {
	Size        uint
	Chromosomes []ChromosomeModel
}

// NewPopulation creates new population instance
func NewPopulation(size uint) *PopulationModel {
	return &PopulationModel{Size: size, Chromosomes: make([]ChromosomeModel, size)}
}

// GenerateChromossomes generates chromossomes of population
func (pop *PopulationModel) GenerateChromossomes(order OrderModel, stores StoresGroup) {
	for k := range pop.Chromosomes {
		pop.Chromosomes[k] = NewChromosome(order.FinalAmount)
	}
}

// GenerateChromossomesGenes randomly generates the genes of each chromossome in population
func (pop *PopulationModel) GenerateChromossomesGenes(order OrderModel, stores StoresGroup) {
	for idx := range pop.Chromosomes {
		pop.Chromosomes[idx].GenerateGenes(order, stores)
	}
}

// GetTopOne returns the chromossome with lower fitness
func (pop *PopulationModel) GetTopOne() ChromosomeModel {
	top := ChromosomeModel{Fitness: math.MaxFloat64}

	for idx := range pop.Chromosomes {
		if pop.Chromosomes[idx].Fitness < top.Fitness {
			top = pop.Chromosomes[idx]
		}
	}

	return top
}

// NaturalSelection randomly selects two chromossos in a population
func (pop *PopulationModel) NaturalSelection() []ChromosomeModel {
	chromo1 := pop.Chromosomes[util.Rand(int(pop.Size))]
	chromo2 := pop.Chromosomes[util.Rand(int(pop.Size))]

	for ; chromo2.Fitness >= chromo1.Fitness; chromo2 = pop.Chromosomes[util.Rand(int(pop.Size))] {
	}

	return []ChromosomeModel{chromo1, chromo2}
}

// ChangeChromossome selects two chromossomes and then replaces
// the worst chromosome with another better
func (pop *PopulationModel) ChangeChromossome(parent1 int, parent2 int, children ChromosomeModel) {
	if pop.Chromosomes[parent1].Fitness > pop.Chromosomes[parent2].Fitness {
		pop.Chromosomes[parent1] = children
	} else {
		pop.Chromosomes[parent2] = children
	}
}

// MakeMutation mutates a chromosome and inserts it into the population
func (pop *PopulationModel) MakeMutation(stores *StoresGroup, mutations int) {
	for i := 0; i < mutations; i++ {
		rand1 := util.Rand(int(pop.Size))
		rand2 := util.Rand(int(pop.Size))

		parent1 := pop.Chromosomes[rand1]
		parent2 := pop.Chromosomes[rand2]

		children := NewGeneration([]ChromosomeModel{
			parent1,
			parent2,
		})

		children.Mutation(config.GetConfig().Mutation.Chance, config.GetConfig().Mutation.Genes, stores)
		children.CalculateFitness()
		pop.ChangeChromossome(rand1, rand2, children)
	}
}
