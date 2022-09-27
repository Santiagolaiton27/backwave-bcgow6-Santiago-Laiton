package main


import(
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)
 

func opMinium(values ... int)(result int){
	var max int = values[0]
	for _, value := range values{
		if value < max{
			max = value
		}
	}
	return max
}

func opAverage(values ... int)(result int){
	var students int
	var sum int
	for _, value := range values{
		sum += value
		students ++
	}
	return sum/students
}
func opMaxium(values ... int)(result int){
	var min int
	for _, value := range values{
		if value > min{
			min = value
		}
	}
	return min
}
func operacionAritmetica(operador string) func(valor1 ... int) int {
	switch operador {
	case "minimum":
		return opMinium
	case "average":
		return opAverage
	case "maximum":
		return opMaxium
	}
	return nil
 }
 

func main(){
	var returnedOperationFunction func(valor1 ... int) int
	returnedOperationFunction = operacionAritmetica(minimum)
	fmt.Println("Minimo : ",returnedOperationFunction(2, 3, 3, 4, 10, 2, 4, 5))
	returnedOperationFunction = operacionAritmetica(average)
	fmt.Println("Promedio : ",returnedOperationFunction(2, 3, 3, 4, 1, 2, 4, 5))
	returnedOperationFunction = operacionAritmetica(maximum)
	fmt.Println("Maximo : ",returnedOperationFunction(2, 3, 3, 4, 1, 2, 4, 15))
	
}