package api

import (
	"time"
)

type Event struct {
	ID     uint64    `json:"id,omitempty"`
	UserID uint64    `json:"user_id,omitempty"`
	Title  string    `json:"title,omitempty"`
	Date   time.Time `json:"date,omitempty"`
}
