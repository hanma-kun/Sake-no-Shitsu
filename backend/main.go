package main

import (
	"log"

	"sake/server"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("[WARNING] Error loading .env, program may use file using default values")
	}
	server.Serve()
}
