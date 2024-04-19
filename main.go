package main

import (
	"github.com/joho/godotenv"
	"levi-api/api/router"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading.env file")
	}
	r := router.Router{}
	r.InitDefault()
}
