package models

type Role struct {
	ID   int
	Name string `gorm:"not null;"`
}
