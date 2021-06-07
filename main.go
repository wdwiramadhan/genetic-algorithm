package main

import (
	"fmt"
	"math"
	"math/rand"
)

type GeneticAlgorithm struct {
	Chromosome [][]float64
	ObjectiveFunction []float64
	Fitness []float64
	Probabilitas []float64
	ProbabilitasCum []float64
}

func Init() *GeneticAlgorithm{
	chromosome := [][]float64{
		{ 12,5,3,8 },
		{ 2,1,8,3 },
		{ 10,4,3,4 },
		{ 20,1,10,6 },
		{ 1,4,3,9 },
		{ 20,5,7,1 },
	}
	return &GeneticAlgorithm{chromosome,[]float64{},[]float64{},[]float64{},[]float64{}}
}


func (ga *GeneticAlgorithm) Evaluation() {
	ga.ObjectiveFunction = []float64{}
	for _,v := range ga.Chromosome {
		objective := math.Abs((v[0] + (v[1]*2) + (v[2]*3) + (v[3]*4)) - 30)
		ga.ObjectiveFunction = append(ga.ObjectiveFunction, objective)
	}
}

func (ga *GeneticAlgorithm) Selection() {
	var totalFitnes float64 = 0
	var probCum float64 = 0
	ga.Fitness = []float64{}
	ga.Probabilitas = []float64{}
	ga.ProbabilitasCum = []float64{}
	for _,v := range ga.ObjectiveFunction {
		result := 1/(v+1)
		totalFitnes += result
		ga.Fitness = append(ga.Fitness, result)
	}
	for _,v := range ga.Fitness {
		result := v/totalFitnes
		probCum += result
		ga.Probabilitas = append(ga.Probabilitas, result)
		ga.ProbabilitasCum = append(ga.ProbabilitasCum, probCum)
	}
}

func (ga *GeneticAlgorithm) Crossover(){
	chromosomeChangePost := []int{0,3,4}
	var chromosomeWillChange [][]float64
	for _, v := range chromosomeChangePost {
		chromosomeWillChange = append(chromosomeWillChange, ga.Chromosome[v])
	}
	i := 0
	for i < len(chromosomeWillChange){
		if i == len(chromosomeWillChange)-1 {
			arr1 := chromosomeWillChange[i][0:2]
			arr2 := chromosomeWillChange[0] [2:len(chromosomeChangePost)+1]
			ga.Chromosome[chromosomeChangePost[i]] = append(arr1,arr2...)
		}else{
			arr1 := chromosomeWillChange[i][0:1]
			arr2 := chromosomeWillChange[i+1] [1:len(chromosomeChangePost)+1]
			ga.Chromosome[chromosomeChangePost[i]] = append(arr1,arr2...)
		}
		i++
	}
}

func (ga *GeneticAlgorithm) Mutation(){
	random1 := rand.Intn(23 - 1) + 1
	random2 := rand.Intn(23 - 1) + 1
	if random1/4 < 4 {
		ga.Chromosome[0][random1%4] = float64(rand.Intn(30 - 1) + 1)
	} else {
		ga.Chromosome[random1/4][random1%4] = float64(rand.Intn(13 - 1) + 1)
	}

	if random2/4 < 4 {
		ga.Chromosome[0][random2%4] = float64(rand.Intn(13 - 1) + 1)
	}else{
		ga.Chromosome[random2/4][random2%4] = float64(rand.Intn(13 - 1) + 1)
	}
}

func (ga *GeneticAlgorithm) MutationCheck() bool {
	for _,v := range ga.Chromosome {
		if v[0] + (v[1]*2) + (v[2]*3) + (v[3]*4) == 30{
			fmt.Println(v)
			return true
		}
	}
	return false
}


func  main()  {
	ga := Init()
	generation := 0
	for {
		ga.Evaluation()
		ga.Selection()
		ga.Crossover()
		ga.Mutation()
		res := ga.MutationCheck()
		if(res){
			fmt.Println(generation, "Generation")
			break
		}
		generation++
	}
}
