package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"` //标签
	Password string `json:"password"`
}

type Post struct {
	gorm.Model
	Title   string
	Content string `gorm:"type:text"`
	Tag     string
}
