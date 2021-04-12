package api

import (
	"encoding/json"
	"net/http"
)

// Login validates user's login credentials
func Login(w http.ResponseWriter, r *http.Request, server *Server) {
	var user User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Error occurred while decoding body.")
		return
	}

	var ok bool
	sqlQueryString := `SELECT 1 FROM users WHERE username=$1 and password=$2`
	server.DB.QueryRow(sqlQueryString, user.Username, user.Password).Scan(&ok)

	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	response := map[string]string{"status": "Accepted"}
	RespondWithJSON(w, http.StatusOK, response)
}

// SearchUsers searches users with a certain prefix
func SearchUsers(w http.ResponseWriter, r *http.Request, server *Server) {
	keys, ok := r.URL.Query()["username"]
	if !ok {
		RespondWithError(w, http.StatusBadRequest, "Payload missing 'username'")
		return
	}

	username := keys[0] + "%"
	sqlQueryString := `SELECT username FROM users WHERE username LIKE $1`
	rows, err := server.DB.Query(sqlQueryString, username)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Error while querying table")
		return
	}

	var users = []string{}
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			RespondWithError(w, http.StatusBadRequest, "Error while parsing query rows")
			return
		}
		if _, ok := server.UserConns[username]; ok {
			users = append(users, username)
		}
	}

	response := map[string][]string{"usernames": users}
	RespondWithJSON(w, http.StatusOK, response)
}
