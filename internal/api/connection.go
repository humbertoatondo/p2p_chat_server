package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// var connections = make(map[string]*websocket.Conn)

// ConnectSocket handles the web socket connections and adds users to exisiting connections
func ConnectSocket(w http.ResponseWriter, r *http.Request, server *Server) {
	keys, ok := r.URL.Query()["username"]
	if !ok {

		return
	}

	username := keys[0]

	ws, err := websocket.Upgrade(w, r, nil, 0, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer ws.Close()

	server.UserConns[username] = ws

	fmt.Printf("[New connection] %s: %s\n", username, r.RemoteAddr)

	for {
		// Receive message
		_, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}

		var objmap map[string]json.RawMessage
		err = json.Unmarshal(message, &objmap)

		var receiver string
		err = json.Unmarshal(objmap["receiver"], &receiver)

		if err := server.UserConns[receiver].WriteMessage(websocket.TextMessage, message); err != nil {
			fmt.Printf("Error sending message: %v\n", err)
		}

		// for _, conn := range server.UserConns {
		// 	if conn == ws {
		// 		continue
		// 	}
		// 	if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
		// 		fmt.Printf("Error: %v\n", err)
		// 	}
		// }

		// fmt.Printf("Message received: %s\n\n\n", message)

		// Reponse messaege

		// if err := ws.WriteMessage(messageType, []byte({"hello friend 123"})); err != nil {
		// 	fmt.Printf("Error: %v\n", err)
		// 	break
		// }
	}

	delete(server.UserConns, username)
}
