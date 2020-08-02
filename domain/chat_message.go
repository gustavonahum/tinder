package domain

import "time"

type ChatMessage struct {
	ID int64 `json:"id"`
	IDSender int64 `json:"id_sender"`
	IDReceiver int64 `json:"id_receiver"`
	Text string `json:"text"`
	Like bool `json:"like"`
	CreatedAt time.Time `json:"created_at"`
}