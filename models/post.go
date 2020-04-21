package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Title  string `gorm:"type:nvarchar(120);not null;"`
	UserId int    `gorm:"not null;"`
	User   *User
}
