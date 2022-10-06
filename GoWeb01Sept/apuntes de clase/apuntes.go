package main

import (
	"fmt"
	"net/http"
)

/*
func main() {
	type product struct {
		Name      string
		Price     float64
		Published bool
	}
	noteBook := product{
		Name:      "MackBook pro",
		Price:     12343.4244,
		Published: true,
	}
	jsonData, erre := json.Marshal(noteBook)

	if erre != nil {
		log.Fatal(erre)
	}
	fmt.Printf("Producto :\n %s", string(jsonData))
}*/

func holaHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hola\n")
}

func main() {
	http.HandleFunc("/hola", holaHandler)
	http.ListenAndServe(":8080", nil)
}
