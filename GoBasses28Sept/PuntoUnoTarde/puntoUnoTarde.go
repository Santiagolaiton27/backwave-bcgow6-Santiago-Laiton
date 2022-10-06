package main

/*
En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: el salario ingresado no alcanza el mínimo imponible" y
lánzalo en caso de que “salary” sea menor a 150.000. Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
*/
import (
	"fmt"
)

type error interface {
	Error()
}
type Error struct {
	message string
}

func (e *Error) Error(err string) string {
	e.message = err
	return e.message
}
func main() {
	salary := 14155454
	//salary := 49000
	MyError := Error{
		"Hola",
	}
	if salary < 150000 {
		panic(MyError.Error("el salario ingresado no alcanza el mínimo imponible"))
	}
	fmt.Println("Tienes que pagar Impuestos")
}
