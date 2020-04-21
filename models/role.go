package models

import (
	"strings"
	"unicode"
)

type Role struct {
	ID          int    `json:"id,omitempty" gorm:"primary_key"`
	Name        string `json:"name,omitempty" gorm:"type:nvarchar(32);not null;"`
	Description string `json:"description,omitempty" gorm:"type:nvarchar(126);"`
	Users       []*User
}

func (r *Role) BeforeSave() {
	r.Name = strings.ToLowerSpecial(unicode.TurkishCase, r.Name)
}
