package model

import "time"

type Item struct {
	ID         string    `json:"id" gorm:"primary_key"`
	Name       string    `json:"group_id"`
	Memo       string    `json:"memo"`
	Price      int       `price:"price"`
	GroupID    string    `group_id:"group_id"`
	HasSettled bool      `has_settled:"has_settled"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
