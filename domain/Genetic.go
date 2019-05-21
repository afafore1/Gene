package domain

import (
	"math/rand"
	"time"
)

type Chromosome struct {
	Gene []interface{}
	Fitness interface{}
}

// GenerateSingleChromosome returns a single chromosome. Take care of this in your code
func (c *Chromosome) GenerateSingleChromosome() interface{} {
	return nil
}

type Population struct {
	Chromosomes []Chromosome
	TopChromosomes []Chromosome
}

func (p *Population) GeneratePopulation(size int, setGene func(c *Chromosome), calculateFitness func(c *Chromosome)) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		c := Chromosome{}
		setGene(&c)
		calculateFitness(&c)
		p.Chromosomes = append(p.Chromosomes, c)
	}
}

func (p *Population) GetRandomParents(bestChromosomes []Chromosome) (Chromosome, Chromosome) {
	rand.Seed(time.Now().UTC().UnixNano()) //no specific reason for making this UTC BTW
	pos1 := rand.Intn(len(bestChromosomes))
	pos2 := rand.Intn(len(bestChromosomes))
	parent1 := bestChromosomes[pos1]
	parent2 := bestChromosomes[pos2]
	return parent1, parent2
}

func (p *Population) GenerateChildren(ChildrenCount int, generateSingleChromosome func() interface{}, calculateFitness func(c *Chromosome)) []Chromosome {
	rand.Seed(time.Now().UTC().UnixNano()) //no specific reason for making this UTC BTW
	var children []Chromosome
	for len(children) < ChildrenCount {
		parent1, parent2 := p.GetRandomParents(p.TopChromosomes)
		//fmt.Println("parent1 and 2 ", parent1, parent2)
		var child Chromosome
		for i := 0; i < len(parent1.Gene); i++ {
			choice := rand.Intn(10) // we want three segments here
			if choice <= 3 {
				child.Gene = append(child.Gene, parent1.Gene[i])
			} else if choice <= 6 {
				child.Gene = append(child.Gene, parent2.Gene[i])
			} else {
				singleChromosome := generateSingleChromosome()
				child.Gene = append(child.Gene, singleChromosome)
			}
		}
		calculateFitness(&child)
		//fmt.Println("child")
		//fmt.Println(child)
		children = append(children, child)
	}
	return children
}