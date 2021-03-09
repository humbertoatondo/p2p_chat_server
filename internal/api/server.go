package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Server is where routers and database are stored
type Server struct {
	Router *mux.Router
	DB     *sql.DB
}

/*
Initialize is the setup function where the server connects to
the database, the router is initialized and routes are created.
*/
func (server *Server) Initialize(host, user, password, dbname string) {
	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		user,
		password,
		dbname,
	)

	var err error
	if server.DB, err = sql.Open("postgres", connectionString); err != nil {
		log.Fatal(err)
	}

	server.Router = mux.NewRouter()
}

// Run is used for the server to start listening
func (server *Server) Run(port string) {
	log.Fatal(http.ListenAndServe(port, server.Router))
}
