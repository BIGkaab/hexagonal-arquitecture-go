package entity

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID          int    `gorm:"not null;unique"`
	Name        string `gorm:"size:50;not null;unique"`
	Description string `gorm:"size:150;not null"`
	Punctuation int    `gorm:"not null"`
}
