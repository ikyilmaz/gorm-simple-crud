package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int     `json:"id,omitempty" gorm:"primary_key"`
	FirstName string  `json:"firstName,omitempty" gorm:"not null;"`
	LastName  string  `json:"lastName,omitempty" gorm:"not null;"`
	Email     string  `json:"email,omitempty" gorm:"not null;"`
	Password  string  `json:"password,omitempty" gorm:"not null;"`
	IsActive  bool    `json:"isActive,omitempty" gorm:"DEFAULT:true"`
	RoleID    int     `json:"roleId,omitempty" gorm:"not null;"`
	Role      Role    `json:"role,omitempty"`
	Posts     []*Post `json:"posts,omitempty"`
	TimeStamps
}

func (u *User) BeforeSave() {
	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(hash)
}
