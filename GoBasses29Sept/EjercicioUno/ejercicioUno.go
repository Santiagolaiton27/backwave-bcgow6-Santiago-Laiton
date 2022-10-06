package main

import (
	"fmt"
	"os"
)

func main() {
	defer func() {
		fmt.Println("Termino la ejecucion")
	}()
	_, erro := os.Open("customer.txt")
	if erro != nil {
		panic("El archivo no se encontro de manera exitosa")
	}
	fmt.Println("Se ejecuto con exito")
}
