package connection

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var connections = make(map[string]*websocket.Conn)

func ConnectionTest(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r, nil, 0, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer ws.Close()

	connections[r.RemoteAddr] = ws

	fmt.Printf("New connection: %s\n", r.RemoteAddr)

	for {
		// Receive message
		_, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			break
		}

		for _, conn := range connections {
			if conn == ws {
				continue
			}
			if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}

		fmt.Printf("Message received: %s\n\n\n", message)

		// Reponse messaege

		// if err := ws.WriteMessage(messageType, []byte({"hello friend 123"})); err != nil {
		// 	fmt.Printf("Error: %v\n", err)
		// 	break
		// }
	}

	delete(connections, r.RemoteAddr)
}
