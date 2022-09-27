package main

import (
	"fmt"
)


func main () {
	var age int
	var salary float64
	var name string
	var antiquity int

	fmt.Println("Cual es tu nombre")
	fmt.Scanln(&name)
	fmt.Println("Cual es tu salario")
	fmt.Scanln(&salary)
	fmt.Println("Cual es tu edad")
	fmt.Scanln(&age)
	fmt.Println("Cuantos años lleva trabajando en la empresa")
	fmt.Scanln(&antiquity)

	if(age < 22){
		fmt.Println("Hola",name,"Lo sentimos no podemos aprobarte el prestamo de ",salary," ya que no cumples con la edad minima")
		return
	}
	if(antiquity<1){
		fmt.Println("Hola",name,"Lo sentimos no podemos aprobarte el prestamo de ",salary," ya que no cumples con la antiguedad minima para poderte aprobar el prestamo")
		return
	}
	fmt.Println("Hola", name,"en hora buena te hemos aprobado tu credito")
}