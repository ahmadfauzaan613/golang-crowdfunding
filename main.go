package main

import (
	"crowdfunding/config"
	"crowdfunding/databases"
	"log"
)

func main() {

	cfg := config.LoadConfig()

	db := databases.DBconnection(cfg)

	if db == nil {
		log.Fatal("Failed to connect DB")
	}

	log.Println("Application started...")
}
