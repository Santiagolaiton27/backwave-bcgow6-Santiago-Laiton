package main

import(
	"fmt"
	"reflect"
)

func main(){
	var(
		humedad int = 36
		temperatura int = 25
		presion float64 = 14815.15
	)

	fmt.Println(humedad)
	fmt.Println(temperatura)
	fmt.Println(presion)

	fmt.Println(reflect.TypeOf(humedad))
	fmt.Println(reflect.TypeOf(temperatura))
	fmt.Println(reflect.TypeOf(presion))
}