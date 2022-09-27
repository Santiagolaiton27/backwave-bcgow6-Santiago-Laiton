package main

import (
	"fmt"
	"errors"
)

func calculateAverage(grades ... float64)(average float64, err error){
	var result float64
	var students int
	for _, value := range grades{
		if(value < 0){
			return 0, errors.New("Las notas no pueden ser negativas") 
		}
		result += value
		students ++
	}
	return result/float64(students),nil
	
}
func main(){
		var average ,err  = calculateAverage(2.6,-4.6,3.4,5.6,3.6,5.4,3.5,5.6,3.5,5.5,3.5,3.5)
		if(err != nil){
			fmt.Println(err)
		}else{
			fmt.Println(average)
		}		
}