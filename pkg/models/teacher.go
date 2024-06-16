package models

import "time"

type Teacher struct {
	Id        uint      `json:"id" gorm:"primary_key;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`

	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Avatar      string `json:"avatar" gorm:"not null"`
}
