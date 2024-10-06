package models

import "time"

type Todo struct {
	TodoID    int       `json:"todo_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
