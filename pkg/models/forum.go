package models

import "time"

// forum thread
// a user creates a thread for a particular course with a title and description
type Thread struct {
	Id        uint      `json:"id" gorm:"primary_key;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`

	UserId   uint   `json:"user_id" gorm:"not null"`
	CourseID uint   `json:"course_id" gorm:"not null"`
	Course   Course `json:"course" gorm:"foreignkey:CourseID"`
	// user who created the thread
	User User `json:"user" gorm:"foreignkey:UserID"`
	// title of the thread
	Title string `json:"title" gorm:"not null"`
	// description of the thread
	Body string `json:"body" gorm:"not null"`
	// comments on the thread
	Comments []Comment `json:"comments" gorm:"foreignkey:ThreadID"`
}

type Comment struct {
	Id        uint      `json:"id" gorm:"primary_key;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`

	UserId   uint   `json:"user_id" gorm:"not null"`
	ThreadID uint   `json:"thread_id" gorm:"not null"`
	Thread   Thread `json:"thread" gorm:"foreignkey:ThreadID"`
	// user who created the comment
	User User `json:"user" gorm:"foreignkey:UserID"`
	// content of the comment
	Body string `json:"body" gorm:"not null"`
}
