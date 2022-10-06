package main

/*
Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,  se implemente “errors.New()”.
*/
import (
	"errors"
	"fmt"
)

func main() {
	//salary := 14155454
	salary := 49000
	if salary < 150000 {
		panic(errors.New("Tu salario no tiene que pagar impuestos"))
	}
	fmt.Println("Tienes que pagar Impuestos")
}
