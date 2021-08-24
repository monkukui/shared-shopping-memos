package model

import "time"

type Group struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"group_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
