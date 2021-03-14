package user

import (
	"fmt"
	"net/http"
)

// Login validates user's login credentials
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("hello"))
}
