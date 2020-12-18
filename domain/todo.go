package domain

import "time"

//Todo Model
type Todo struct {
	ID          int       `json:"id"`
	User        *User     `json:"user"`
	TaskName    string    `json:"name"`
	Description string    `json:"description"`
	Created     time.Time `json:"creation_date"`
	Updated     time.Time `json:"completion_date"`
	Completed   bool      `json:"status"`
}
