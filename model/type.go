package model

import (
	"time"
)

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type User struct {
	Model
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type Token struct {
	Model
	UserID uint
	Body   string
}

type Articles struct {
	Model
	UserID	uint
	Title	string `json:"title" building"required"`
	body	string `json:"body" building"required"`
}