package main



import (
	"fmt"
	"errors"
)
func calculateTax(salary float64)(salaryFinal float64, err error){
	if(salary < 0){
		return 0, errors.New("El salario no puede ser negativo")
	}
	if(salary < 150_000 && salary > 50_000){
		return salary * 0.17 , nil
	}
	if(salary > 150_000){
		return salary * (0.17+0.1),nil
	}
	return 0, nil
}

func main(){

	fmt.Println("Ingrese cual es su salario")
	var salary float64
	fmt.Scanln(&salary)
	tax ,error := calculateTax(salary)
	if(error != nil){
		fmt.Println(error)
	}else{
		fmt.Println("Los impustos fueron : ",tax)
	}
}