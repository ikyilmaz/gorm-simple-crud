package models

type Post struct {
	ID     int    `json:"id,omitempty" gorm:"primary_key"`
	Title  string `json:"title,omitempty" gorm:"type:nvarchar(120);not null;"`
	UserID int    `json:"userId,omitempty" gorm:"not null;"`
	User   *User  `json:"user,omitempty"`
	TimeStamps
}
