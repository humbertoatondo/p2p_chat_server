package api

import (
	"database/sql"

	_ "github.com/lib/pq" // This package is needed for the postgres database

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// Server is where routers and database are stored
type Server struct {
	Router    *mux.Router
	DB        *sql.DB
	UserConns map[string]*websocket.Conn
}

// User is the structure for the users table
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type WebRtcSdp struct {
	Type     string `json:"type"`
	Receiver string `json:"receiver"`
	Sdp      string `json:"sdp"`
}
