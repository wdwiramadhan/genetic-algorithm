package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)


type GeneticAlgorithm struct {
	Chromosome [][]float64
	ObjectiveFunction []float64
	Fitness []float64
	Probabilitas []float64
	ProbabilitasCum []float64
}

func Init() *GeneticAlgorithm{
	rand.Seed(time.Now().UnixNano())
	var chromosome [][]float64
	i := 0
	for i < 6 {
		j :=0
		var arr []float64
		for j < 4 {
			arr = append(arr, float64(rand.Intn(30 - 0) + 1))
			j++
		}
		chromosome = append(chromosome, arr)
		i++
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

	var Rouletes []float64
	j := 0
	for j < 6 {
		Rouletes = append(Rouletes, 0+rand.Float64() * (1-0))
		j++
	}
	
	newChromosome := ga.Chromosome
	for indexRoulete, roulete := range Rouletes {
		for indexProbab, probab := range ga.ProbabilitasCum{
			if roulete > probab && indexProbab < 5 && roulete < ga.ProbabilitasCum[indexProbab+1] {
				newChromosome[indexRoulete] = ga.Chromosome[indexProbab]
			}
		}
	}
	ga.Chromosome = newChromosome
}

func (ga *GeneticAlgorithm) Crossover(){
	rand.Seed(time.Now().UnixNano())
	chromosomeChangePost := []int{rand.Intn(5 - 0) + 0,rand.Intn(5 - 0) + 0,rand.Intn(5 - 0) + 0}
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
	rand.Seed(time.Now().UnixNano())
	random1 := rand.Intn(24 - 1) + 1
	random2 := (rand.Intn(24 - 1) + 1)/10
	if random1/4 < 4 {
		ga.Chromosome[0][random1%4] = float64(rand.Intn(30 - 0) + 0)
	} else {
		ga.Chromosome[random1/4][random1%4] = float64(rand.Intn(30 - 0) + 0)
	}

	if random2/4 < 4 {
		ga.Chromosome[0][random2%4] = float64(rand.Intn(30 - 0) + 0)
	}else{
		ga.Chromosome[random2/4][random2%4] = float64(rand.Intn(30 - 0) + 0)
	}
}

func (ga *GeneticAlgorithm) MutationCheck() bool {
	for _,v := range ga.Chromosome {
		if v[0] + (v[1]*2) + (v[2]*3) + (v[3]*4) == 30{
			fmt.Println("\nBest Chromosome",v)
			return true
		}
	}
	return false
}


func  main()  {
	ga := Init()
	fmt.Println("Inisialiasi")
	fmt.Println(ga.Chromosome)
	fmt.Println("Evaluation")

	generation := 0
	for {
		ga.Evaluation()
		fmt.Println(ga.ObjectiveFunction)
	
		fmt.Println("Selection")
		ga.Selection()
		fmt.Println("Fitnes", ga.Fitness)
		fmt.Println("Probability",ga.Probabilitas)
		fmt.Println("Cumulative Probability",ga.ProbabilitasCum)
		fmt.Println("After Selection",ga.Chromosome)
	
		ga.Crossover()
		fmt.Println("Crossover", ga.Chromosome)
		ga.Mutation()
		fmt.Println("Mutation", ga.Chromosome)
		res := ga.MutationCheck()
		if(res){
			fmt.Println(generation, "Generation")
			break
		}
		fmt.Println()
		generation++
	}
}
