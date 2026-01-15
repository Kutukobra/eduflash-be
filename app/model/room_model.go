package model

import "time"

type Room struct {
	ID         string    `json:"id"`
	Room_Name  string    `json:"room_name"`
	Created_At time.Time `json:"created_at"`
	Owner_ID   string    `json:"owner_id"`
}

type RoomStudent struct {
	Room_ID      string `json:"room_id"`
	Student_Name string `json:"student_name"`
}
