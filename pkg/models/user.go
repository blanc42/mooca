package models

import "time"

type User struct {
	// gorm.Model
	Id        uint      `json:"id" gorm:"primary_key;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`

	Username string `json:"username" gorm:"not null;unique"`
	Email    string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
	Role     Role   `json:"role" gorm:"not null;foreignkey:RoleId"`
	RoleId   uint   `json:"role_id" gorm:"not null"`
}

type Role struct {
	Id        uint      `json:"id" gorm:"primary_key;not null"`
	CreatedAt time.Time `json:"-" gorm:"not null"`
	UpdatedAt time.Time `json:"-" gorm:"not null"`

	Name string `json:"name" gorm:"not null;unique"`
}
