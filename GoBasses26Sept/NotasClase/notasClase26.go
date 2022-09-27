package main



import(
	"fmt"
)
const(
	operadorSuma = "+"
	operadorResta = "-"
	operadorMulti = "*"
	operadorDivi = "/"
)
func inspectionVariable(numero int){
	if numero>0{
		fmt.Println("A es positivo")
	}else if numero < 0{
		fmt.Println("A es negativo")
	}else {
		fmt.Println("A es cero")
	}
}
func suma(a,b float64)float64{
	resultado := a+b
	return resultado
}

func operationArimetic(a,b float64, operator string)float64{
	var result float64
	switch operator{
	case operadorSuma:
		result = a+b
	case operadorResta:
		result = a-b
	case operadorMulti:
		result = a*b
	case operadorDivi:
		if(b!=0){
			result = a/b
		}
	}
	return result
}
func main(){
	/*
	a,b,c,d := 0,2,-1,3
	inspectionVariable(a)
	inspectionVariable(b)
	inspectionVariable(c)
	inspectionVariable(d)

	suma := suma(1.3,4.3)

	fmt.Println(suma)
	*/
	fmt.Println(operationArimetic(1.2,3.4,operadorMulti))
	fmt.Println(operationArimetic(1.2,0,operadorDivi))
	fmt.Println(operationArimetic(1.2,3.4,operadorSuma))
	fmt.Println(operationArimetic(1.2,3.4,operadorResta))
}