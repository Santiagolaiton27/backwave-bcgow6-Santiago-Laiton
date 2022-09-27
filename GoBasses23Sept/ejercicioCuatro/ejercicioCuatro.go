package main

import(
	"fmt"
)

func main(){
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	benjamin := employees["Benjamin"]
	fmt.Println(benjamin)
	var repeat int
	for key, element := range employees{
		if(element > 21){
			repeat++
		}
		fmt.Println("Key => ",key,"Element =>",element)
	}
	employees["Federico"]=25
	delete(employees, "Pedro")
	fmt.Println(employees)
}