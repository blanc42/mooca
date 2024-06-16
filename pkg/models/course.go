package models

import "time"

type Course struct {
	Id        uint      `json:"id" gorm:"primary_key;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`

	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	// institutes
	// teachers
	// tags
}

type Tag struct {
	Id        uint      `json:"id" gorm:"primary_key;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`

	Name     string `json:"name" gorm:"not null"`
	Approved bool   `json:"approved" gorm:"not null"`
	// courses
}
