package main

import(
	"fmt"
)

type Matrix struct{
	Listas [][] int
	Max int
	Tall int
	Width int
	Square  bool

}

func(matrix *Matrix)Set(nueva [][]int){
	matrix.Listas = nueva
}

func(matrix *Matrix)FindMaxAndMetrics(){
	var max,tall,width int = -1000000,0,0
	
	for i := 0; i < len(matrix.Listas); i++ {
		for j := 0; j < len(matrix.Listas[i]); j++ {
			if(matrix.Listas[i][j]>max){
				max = matrix.Listas[i][j]
			}
		width = j
		}
	tall = i
	}
	matrix.Max = max
	if(tall == width){
		matrix.Square = true
	}
	matrix.Tall = tall
	matrix.Width = width
}

func main(){
	lista := [][]int{{1, 50, 15_0000, 1}, {1, 50, 15_000, 1}, {1, 50, 15_000, 1}, {1, 50, 15_000, 1}}
	matrix := Matrix{

		// Listas: [][]int{
		// {1, 50, 50000, 1}, 
		// {1, 50, 20000, 1}, 
		// {1, 50, 10000, 1}, 
		// {1, 50, 30000, 1},
		// },
	}
	matrix.Set(lista)
	for i:=0 ;i<len(matrix.Listas); i++ {
		fmt.Println(matrix.Listas[i])
	}
	matrix.FindMaxAndMetrics()
	fmt.Println(matrix)
	
}