package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fullname string `json:"fullname"`
	Email    string `json:"email" gorm:"unique not null"`
	Password string `json:"password"`
}
