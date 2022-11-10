package db

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func ConnectDatabase() (engine *gin.Engine, db *sql.DB) {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error :Loading.env")
	// }
	configDB := mysql.Config{
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "storage",
	}
	db, err := sql.Open("mysql", configDB.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	engine = gin.Default()
	return engine, db
}
