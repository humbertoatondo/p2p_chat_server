package models

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// Server is where routers and database are stored
type Server struct {
	Router    *mux.Router
	DB        *sql.DB
	UserConns map[string]*websocket.Conn
}

type fn func(http.ResponseWriter, *http.Request)
