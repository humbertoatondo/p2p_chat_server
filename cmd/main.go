package main

import (
	"log"
	"os"

	"github.com/humbertoatondo/p2p_chat_server/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading dotenv file")
	}

	server := api.Server{}
	server.Initialize(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	server.Run(":3000")
}
