package user

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type user struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Login validates user's login credentials
func Login(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var usr user
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&usr); err != nil {
		w.Write([]byte("Error occurred while decoding body.\n"))
		return
	}

	var isValidUser bool
	sqlQueryString := `SELECT 1 FROM users WHERE name=$1 and password=$2`
	db.QueryRow(sqlQueryString, usr.Name, usr.Password).Scan(&isValidUser)

	if !isValidUser {
		respondWithError(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	response := map[string]string{"status": "Accepted"}
	respondWithJSON(w, http.StatusOK, response)
}

func SearchUsers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	keys, ok := r.URL.Query()["searchTerm"]
	if !ok {
		respondWithError(w, http.StatusBadRequest, "Payload missing 'searchTerm'")
		return
	}

	searchTerm := keys[0] + "%"
	sqlQueryString := `SELECT name FROM users WHERE name LIKE $1`
	rows, err := db.Query(sqlQueryString, searchTerm)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error while querying table")
		return
	}

	var users = []string{}
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			respondWithError(w, http.StatusBadRequest, "Error while parsing query rows")
			return
		}
		users = append(users, username)
	}

	response := map[string][]string{"usernames": users}
	respondWithJSON(w, http.StatusOK, response)
}
