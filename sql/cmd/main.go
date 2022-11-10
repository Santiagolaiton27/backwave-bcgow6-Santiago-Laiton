package main

import (
	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/sql/cmd/server/routes"
	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/sql/pkg/db"
)

func main() {
	engine, db := db.ConnectDatabase()
	router := routes.NewRouter(engine, db)
	router.MapRoutes()

	engine.Run(":8080")

}
