package main



import (
	"fmt"
	"errors"
)

var categories = map[string]float64{
        "C": 1.000,
        "B": 1.500,
        "A": 3.000,
	}
	

func calculatorSalary(category string, workHour float64)(salary float64, err error){
	var percentage float64
	switch category{
	case "A":
		percentage := workHour * categories["A"]*0.5
		return (workHour * categories["A"])+percentage,nil
	case "B":
		percentage := workHour * categories["B"]*0.2
		return (workHour * categories["B"])+percentage,nil
	case "C":
	return (workHour * categories["C"])+percentage,nil
	}
	return 0, errors.New("La categoria que ingreso no existe")

}

func main(){
	var workHour float64
	var category string
	fmt.Println("Ingrese cuantas horas trabajo el empleado en el mes")
	fmt.Scanln(&workHour)
	fmt.Println("Ingrese la categoria del empleado")
	fmt.Scanln(&category)
	salary ,err := calculatorSalary(category,workHour)
	if(err != nil){
		fmt.Println(err)	
	}else{
		fmt.Println(salary)		
	}
	





}