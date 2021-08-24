package model

import "time"

const (
	ColorRed = iota
	ColorBlue
	ColorGreen
	ColorPink
	ColorYellow
)

type User struct {
	ID        string    `json:"id" gorm:"primary_key"`
	GroupID   string    `json:"group_id"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Color     int       `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
