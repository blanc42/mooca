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
	// add more fields
	Institutes []Institute `json:"institutes" gorm:"many2many:course_institutes;"`
	Teachers   []Teacher   `json:"teachers" gorm:"many2many:course_teachers;"`
	Tags       []Tag       `json:"tags" gorm:"many2many:course_tags;"`
	Duration   int         `json:"duration" gorm:"not null"` // Duration in hours
	Credits    float64     `json:"credits" gorm:"not null"`  // Course credit points
	Level      string      `json:"level" gorm:"not null"`    // e.g., Beginner, Intermediate, Advanced
}

type Tag struct {
	Id        uint      `json:"id" gorm:"primary_key;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`

	Name     string `json:"name" gorm:"not null"`
	Approved bool   `json:"approved" gorm:"not null"`
	// courses
}
