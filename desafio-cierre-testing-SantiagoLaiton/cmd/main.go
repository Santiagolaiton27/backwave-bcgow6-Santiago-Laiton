package main

import (
	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/desafio-cierre-testing-SantiagoLaiton/cmd/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.MapRoutes(r)

	err := r.Run(":18085")
	if err != nil {
		panic(err.Error())
	}
}
