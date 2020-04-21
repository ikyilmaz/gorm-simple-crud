package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseMode
	FirstName string `gorm:"not null;"`
	LastName  string `gorm:"not null;"`
	Email     string `gorm:"not null;"`
	Password  string `gorm:"not null;"`
	IsActive  bool   `gorm:"DEFAULT:true"`
	RoleID    int    `gorm:"not null;"`
	Role      *Role
	Posts     []*Post
}

func (u *User) BeforeSave() {
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hash)
}

func DefaultUserScope(db *gorm.DB) *gorm.DB {
	return db.
		Select("id, first_name, last_name").
		Where("is_active = ?", true).
		Order("created_at", false)
}
