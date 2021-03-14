package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/humbertoatondo/p2p_chat_server/internal/api/connection"
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
		fmt.Printf("Error: %v\n", err)
	}

	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

// Run is used for the server to start listening
func (server *Server) Run(port string) {
	log.Fatal(http.ListenAndServe(port, server.Router))
}

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/connection", connection.ConnectionTest).Methods("GET")
	server.Router.HandleFunc("/user/login", user)
}
