package model

import "time"

type Room struct {
	ID        string    `json:"id"`
	RoomName  string    `json:"room_name"`
	CreatedAt time.Time `json:"created_at"`
	OwnerId   string    `json:"owner_id"`
}
