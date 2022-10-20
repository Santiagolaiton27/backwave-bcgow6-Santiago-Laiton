package main

import (
	"fmt"

	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/GoTestinUno/pkg/calculador"
)

func main() {
	a, b := 10, 5
	sum := calculador.Add(a, b)
	fmt.Println(sum)
}
