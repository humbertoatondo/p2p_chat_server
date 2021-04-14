package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
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

	ensureUsersTableExists()
	code := m.Run()
	clearUsersTable()
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
func ensureUsersTableExists() {
	if _, err := server.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

// Make sure users table is empty.
func clearUsersTable() {
	server.DB.Exec("DELETE FROM users")
	server.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
}

// Execute http request.
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	server.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

// Ping database to check that there is a connection.
func TestDatabaseConnection(t *testing.T) {
	err := server.DB.Ping()
	if err != nil {
		t.Errorf("Could not connect with postgres database: %v\n", err)
	}
}

func TestUserLogin(t *testing.T) {
	clearUsersTable()

	// Insert row in users table.
	server.DB.Exec("INSERT INTO users (username, password) VALUES ('john_doe', 'password')")

	// Make http post request for user's login.
	var jsonStr = []byte(`{"username":"john_doe", "password": "password"}`)
	req, _ := http.NewRequest("POST", "/user/login", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	// Verify response code is the expected result.
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}
