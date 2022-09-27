package main

import(
	"fmt"
)
const (
	dog = "dog"
	cat = "cat"
)

func dogFood(number int) int{
	return number * 10
}
func catFood(number int) int{
	return number * 5
}
func getAnimal(animal string) func (number int) int{
	switch animal{
	case "dog":
		return dogFood
	case "cat":
		return catFood
	}
	return nil
}

func main(){
	var amount int
	var returnedAnimal func(int) int
	returnedAnimal = getAnimal(dog)
	amount += returnedAnimal(5)
	fmt.Println("comida para perros : ",returnedAnimal(5))
	returnedAnimal = getAnimal(cat)
	amount += returnedAnimal(5)
	fmt.Println("comida para gatos : ",returnedAnimal(10))
	fmt.Println("total de comida : ",amount)
}
