package models

import "time"

type Book struct {
	ID        uint   `gorm:"primary_key"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	ID         uint `gorm:"primary_key"`
	First_Name string
	Last_Name  string
}
