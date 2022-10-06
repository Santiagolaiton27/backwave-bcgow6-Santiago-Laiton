package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func opMinium(values ...int) (result int) {
	var max int = values[0]
	for _, value := range values {
		if value < max {
			max = value
		}
	}
	return max
}

func opAverage(values ...int) (result int) {
	var students int
	var sum int
	for _, value := range values {
		sum += value
		students++
	}
	return sum / students
}
func opMaxium(values ...int) (result int) {
	var min int
	for _, value := range values {
		if value > min {
			min = value
		}
	}
	return min
}
func operacionAritmetica(operador string) (funcion func(valor1 ...int) int, err error) {
	switch operador {
	case "minimum":
		funcion = opMinium
	case "average":
		funcion = opAverage
	case "maximum":
		funcion = opMaxium
	default:
		err = errors.New("Esta funcion no esta definida")
	}
	return
}

func main() {
	var returnedOperationFunction func(valor1 ...int) int
	returnedOperationFunction, err := operacionAritmetica(minimum)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Minimo : ", returnedOperationFunction(2, 3, 3, 4, 10, 2, 4, 5))
	}
	returnedOperationFunction, err = operacionAritmetica(average)
	fmt.Println("Promedio : ", returnedOperationFunction(2, 3, 3, 4, 1, 2, 4, 5))
	returnedOperationFunction, err = operacionAritmetica(maximum)
	fmt.Println("Maximo : ", returnedOperationFunction(2, 3, 3, 4, 1, 2, 4, 15))

}
