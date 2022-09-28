package main

import (
	"fmt"
	"os"
)

func main() {
	str := "ID;PRODUCTO;PRECIO;CANTIDAD\n"
	lista := [4][4]int{{1, 50, 15_000, 1}, {1, 50, 15_000, 1}, {1, 50, 15_000, 1}, {1, 50, 15_000, 1}}
	var producto string
	for i := 0; i < len(lista); i++ {
		for j := 0; j < len(lista[i]); j++ {
			//fmt.Println(lista[i][j])
			producto = fmt.Sprintf("%v,%v,%v,%v\n", lista[i][j], lista[i][j], lista[i][j], lista[i][j])
		}
		fmt.Println(producto)
		str += producto
	}
	doc := []byte(str)
	_ = os.WriteFile("./miFile.csv", doc, 0644)
}
