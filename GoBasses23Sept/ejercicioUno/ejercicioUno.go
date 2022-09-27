package main
import(
	"fmt"
)

func main(){
	var word string
	fmt.Println("Que palabra quiere evaluar")
	fmt.Scanln(&word)
	fmt.Println("El tamaño de la palabra",word," es : ", len(word))
	for _,letter:= range word{
		fmt.Println(string(letter))
	}
}