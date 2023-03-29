package main

import (
	"entlearning/pkg/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	dbManager, err := db.NewDBManager("localhost", 5432, "warddb", "ward", "ward")
	if err != nil {
		log.Fatalln(err)
	}
	err = dbManager.AutoMigrate()
	if err != nil {
		log.Fatalln(err)
	}
	r := gin.Default()
	r.Run()
}
