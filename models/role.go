package models

import (
	"strings"
	"unicode"
)

type Role struct {
	ID          int
	Name        string `gorm:"type:nvarchar(32);not null;"`
	Description string `gorm:"type:nvarchar(126);"`
	Users       []*User
}

func (r *Role) BeforeSave() {
	r.Name = strings.ToLowerSpecial(unicode.TurkishCase, r.Name)
}
