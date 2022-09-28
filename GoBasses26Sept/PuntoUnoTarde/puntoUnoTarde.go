package main

import (
	"fmt"
)

type Persona struct{
	Nombre string 	
	Apellido string	
	Dni string		
	Fecha string	
}
func (student *Persona)CrearPersona(nombre,apellido,dni,fecha string){
	student.Nombre = nombre
	student.Apellido = apellido
	student.Dni = dni
	student.Fecha = fecha
}

func NewStudent() Persona{
	return Persona{}
}

func main() {
	var nombre string
	var apellido string
	var dni string
	var fecha string

	fmt.Println("Ingrese su nombre")
	fmt.Scanln(&nombre)
	fmt.Println("Ingrese su apellido")
	fmt.Scanln(&apellido)
	fmt.Println("Ingrese su documento")
	fmt.Scanln(&dni)
	fmt.Println("Ingrese su feche de nacimiento")
	fmt.Scanln(&fecha)
	
	student := Persona{
		// Nombre :nombre,
		// Apellido:apellido,
		// Dni :dni,
		// Fecha :fecha,
	}

	student.CrearPersona(nombre,apellido,dni,fecha)
	fmt.Println(student)
	// a := NewStudent()
	// fmt.Println(a)
	// a.Nombre = nombre
	// a.Apellido = apellido
	// a.Dni = dni
	// a.Fecha = fecha
	// fmt.Println(a)

}


