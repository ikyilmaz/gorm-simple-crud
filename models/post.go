package models

type Post struct {
	BaseMode
	Title  string `gorm:"type:nvarchar(120);not null;"`
	UserID int    `gorm:"not null;"`
	User   *User
}
