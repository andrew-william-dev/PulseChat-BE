package models

import "time"

type Message struct {
    ID         int       `json:"id"`
    SenderID   int       `json:"sender_id"`
    ReceiverID int       `json:"to"`
    Message    string    `json:"content"`
    CreatedAt  time.Time `json:"created_at"`
}
