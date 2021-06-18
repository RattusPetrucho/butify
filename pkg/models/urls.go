package models

import "time"

type Url struct {
	Id        uint64    `json:"id"`
	Small     string    `json:"small"`
	Origin    string    `json:"origin"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdateAt  time.Time `json:"updated_at" db:"updated_at"`
}
