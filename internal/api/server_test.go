package api

import (
	"os"
	"testing"
)

var server Server

func TestInitialize(t *testing.T) {

	server.Initialize(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

}
