package main

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading dotenv file")
	}

	code := m.Run()
	os.Exit(code)
}
