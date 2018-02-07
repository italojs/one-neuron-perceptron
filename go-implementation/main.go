package main

import (
	"fmt"
)


type Neuron struct{
	Samples [][]float64
	Classes []float64
	Weights []float64
    TrainingRate float64
}

func stepFuction(sum float64)(result int){
	if sum >= 1{
		return 1
	}
	return 0
}

func calculateClass(sample []float64, weights []float64)(result int){
	var sum float64
	for i := range sample{
		sum += sample[i]*weights[i]
	}
	return stepFuction(sum)
}

func train(n *Neuron){
	errors := 1.0
	times := 0
	for errors != 0  {
		errors = 0
		times ++
		for i, class := range (*n).Classes{
			neuronResult := calculateClass((*n).Samples[i],(*n).Weights)
			err := class - float64(neuronResult)
			errors += err
			for j, weight := range (*n).Weights{
				(*n).Weights[j] = weight + ((*n).TrainingRate * (*n).Samples[i][j] * err)
				//fmt.Println("peso atualizado de ", weight, " para ", (*n).Weights[j])
			}
			//fmt.Println("erros: ", err)
		}
	}
}

func main (){
	n := Neuron{
			Samples:[][]float64{{0,0},{0,1}, {1,0},{1,1}},
			Classes:[]float64{0,1,1,1},
			Weights:[]float64{0,0},
			TrainingRate: 0.1}
	train(&n)
	fmt.Println(calculateClass(n.Samples[0], n.Weights))
	fmt.Println(calculateClass(n.Samples[1], n.Weights))
	fmt.Println(calculateClass(n.Samples[2], n.Weights))
	fmt.Println(calculateClass(n.Samples[3], n.Weights))
}