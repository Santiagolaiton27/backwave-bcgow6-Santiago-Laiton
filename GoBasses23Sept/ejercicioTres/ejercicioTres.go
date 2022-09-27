package main

import(
	"fmt"
	"errors"
)
func whatMonth(mes int) (name string, err error){

	switch mes{
	case 1:
		return "Enero", nil
	case 2:
		return "Febrero", nil
	case 3:
		return "Marzo", nil
	case 4:
		return "Abril", nil
	case 5:
		return "Mayo", nil
	case 6:
		return "Junio", nil
	case 7:
		return "Julio", nil
	case 8:
		return "Agosto", nil
	case 9:
		return "Septiembre", nil
	case 10:
		return "Octubre", nil
	case 11:
		return "Noviembre", nil
	case 12:
		return "Dicimebre", nil
	}
	return "", errors.New("Este mes no existe")
	
}
func main(){
	var mes int
	fmt.Println("Cual es numero del mes que quieres saber el nombre")
	fmt.Scanln(&mes)
	name,err := whatMonth(mes)
	if(err != nil){
		fmt.Println(err)
	}else{
		fmt.Println(name)
	}
}