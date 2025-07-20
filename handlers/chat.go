package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	
	"chatapp/utils"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[int]*websocket.Conn)

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Recovered from panic in WebSocketHandler:", err)
		}
	}()

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "userID not found in context", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	clients[userID] = conn
	log.Printf("User %d connected via WebSocket\n", userID)

	for {
		var msg struct {
			To      int    `json:"to"`
			Content string `json:"content"`
		}

		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Read error:", err)
			delete(clients, userID)
			break
		}

		// Save to DB
		_, dbErr := utils.DB.Exec(
			"INSERT INTO messages (sender_id, receiver_id, content) VALUES ($1, $2, $3)",
			userID, msg.To, msg.Content,
		)
		if dbErr != nil {
			log.Println("DB insert error:", dbErr)
			continue
		}

		// Construct broadcastable message
		outMsg := map[string]interface{}{
			"from":    userID,
			"to":      msg.To,
			"content": msg.Content,
		}

		// Send to receiver
		if receiverConn, ok := clients[msg.To]; ok {
			receiverConn.WriteJSON(outMsg)
		}

		// Echo back to sender
		if senderConn, ok := clients[userID]; ok {
			senderConn.WriteJSON(outMsg)
		}
	}
}

