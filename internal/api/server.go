package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type handlerWrapper func(http.ResponseWriter, *http.Request, *Server)

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

	server.UserConns = make(map[string]*websocket.Conn)

	server.initializeRoutes()
}

// Run is used for the server to start listening
func (server *Server) Run(port string) {
	log.Fatal(http.ListenAndServe(port, server.Router))
}

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/connect", server.wrapper(ConnectSocket)).Methods("GET")
	server.Router.HandleFunc("/user/login", server.wrapper(Login)).Methods("POST")
	server.Router.HandleFunc("/searchUsers", server.wrapper(SearchUsers)).Methods("GET")
}

func (server *Server) wrapper(wrapper handlerWrapper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		wrapper(w, r, server)
	}
}
