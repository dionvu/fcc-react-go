package main

import (
	"log"

	"github.com/joho/godotenv"
)

func config() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}
