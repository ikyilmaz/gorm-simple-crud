package models

import (
	"github.com/jinzhu/gorm"
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
	Role      *Role   `json:"role,omitempty"`
	Posts     []*Post `json:"posts,omitempty"`
	TimeStamps
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
