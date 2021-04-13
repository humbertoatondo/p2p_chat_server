package main

import (
	"log"
	"os"
	"testing"

	"github.com/humbertoatondo/p2p_chat_server/internal/api"
	"github.com/joho/godotenv"
)

var server api.Server

func TestMain(m *testing.M) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading dotenv file")
	}

	server.Initialize(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	ensureTableExists()

	code := m.Run()
	os.Exit(code)
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS users
(
    id SERIAL,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(50) NOT NULL,
    CONSTRAINT user_pkey PRIMARY KEY (id)
)`

// Create users table if it doesn't exist.
func ensureTableExists() {
	if _, err := server.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

// Ping database to check that there is a connection.
func TestDatabaseConnection(t *testing.T) {
	err := server.DB.Ping()
	if err != nil {
		t.Errorf("Could not connect with postgres database: %v\n", err)
	}
}
