package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "chatapp/models" 
    "chatapp/utils"
)

func GetMessages(w http.ResponseWriter, r *http.Request) {
    userID, ok := r.Context().Value("userID").(int)
	if !ok {
		http.Error(w, "userID not found in context", http.StatusUnauthorized)
		return
	}
    senderID := userID
    vars := mux.Vars(r)
    receiverID, err := strconv.Atoi(vars["userId"])
    if err != nil {
        http.Error(w, "Invalid receiver ID", http.StatusBadRequest)
        return
    }

    rows, err := utils.DB.Query(`
        SELECT id, sender_id, receiver_id, content, created_at 
        FROM messages 
        WHERE (sender_id = $1 AND receiver_id = $2) OR (sender_id = $2 AND receiver_id = $1)
        ORDER BY created_at ASC`, senderID, receiverID)
    if err != nil {
        http.Error(w, "Error fetching messages", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var messages []models.Message
    for rows.Next() {
        var msg models.Message
        if err := rows.Scan(&msg.ID, &msg.SenderID, &msg.ReceiverID, &msg.Message, &msg.CreatedAt); err != nil {
            http.Error(w, "Error scanning message", http.StatusInternalServerError)
            return
        }
        messages = append(messages, msg)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(messages)
}