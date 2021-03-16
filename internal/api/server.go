package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/humbertoatondo/p2p_chat_server/internal/api/connection"
	"github.com/humbertoatondo/p2p_chat_server/internal/api/user"
)

// Server is where routers and database are stored
type Server struct {
	Router *mux.Router
	DB     *sql.DB
}

type fn func(http.ResponseWriter, *http.Request, *sql.DB)

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
	server.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		fmt.Printf("Error: %v\n", err)
		return
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
	server.Router.HandleFunc("/user/login", server.wrapper(user.Login)).Methods("POST")
	server.Router.HandleFunc("/user/searchUsers", server.wrapper(user.SearchUsers)).Methods("GET")
}

func (server *Server) wrapper(f fn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r, server.DB)
	}
}
