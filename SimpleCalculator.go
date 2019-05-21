package main

import (
	"P/domain"
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"time"
)

const target = 30

func main() {
	rand.Seed(time.Now().UnixNano())
	p := domain.Population{}
	populationSize := 100
	topChromosomeCount := 30
	p.GeneratePopulation(populationSize, setGene, calculateFitness)
	//termination points
	for i := 0; i < 100; i++ { //max iteration
		fmt.Println("Iteration ", i)
		p.TopChromosomes = p.TopChromosomes[:0]
		getTopXChromosomes(&p, topChromosomeCount)
		fmt.Println(p.TopChromosomes)
		children := p.GenerateChildren(len(p.Chromosomes)-topChromosomeCount, generateSingleChromosome, calculateFitness)
		topChromosome := p.TopChromosomes[0]
		topResult := calculateResult(&topChromosome)
		fmt.Printf("Top chromosome has fitness of %d and value of %d:", topChromosome.Fitness, topResult)
		if topResult == target { // another termination point
			fmt.Println("Top Chromosome has gene: ", topChromosome.Gene)
			break
		}
		// kill off bad parents
		p.Chromosomes = p.Chromosomes[:0]
		p.Chromosomes = append(p.Chromosomes, p.TopChromosomes...)
		p.Chromosomes = append(p.Chromosomes, children...)
	}
}

func getTopXChromosomes(p *domain.Population, topChromosomeCount int) {
	chromosomeMap := make(map[int] domain.Chromosome)
	var keys []int
	for _, c := range p.Chromosomes {
		keys = append(keys, getInt(c.Fitness))
		chromosomeMap[getInt(c.Fitness)] = c
	}
	sort.Ints(keys)
	for _, k := range keys {
		if len(p.TopChromosomes) == topChromosomeCount {
			break
		}
		p.TopChromosomes = append(p.TopChromosomes, chromosomeMap[k])
	}
}

func generateSingleChromosome() interface{} {
	return rand.Intn(target)
}

func setGene(c *domain.Chromosome) {
	for i := 0; i < 4; i++ {
		c.Gene = append(c.Gene, generateSingleChromosome())
	}
}

// 1a + 2b + 3c + 4c = target
func calculateFitness(chromosome *domain.Chromosome) {
	diff := Abs(target - calculateResult(chromosome))
	chromosome.Fitness = diff / 1
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func calculateResult(chromosome *domain.Chromosome) int {
	a, b, c, d := chromosome.Gene[0], chromosome.Gene[1], chromosome.Gene[2], chromosome.Gene[3]
	return getIntValue(1, a) + getIntValue(2, b) + getIntValue(3, c) + getIntValue(4, d)
}

func getIntValue(n int, d interface{}) int {
	var intType = reflect.TypeOf(int64(0))
	v := reflect.ValueOf(d)
	v = reflect.Indirect(v)
	intValue := v.Convert(intType)
	return int(int64(n) * intValue.Int())
}

func getInt(d interface{}) int {
	var intType = reflect.TypeOf(int64(0))
	v := reflect.ValueOf(d)
	v = reflect.Indirect(v)
	intValue := v.Convert(intType)
	return int(intValue.Int())
}
