package models

import (
	"github.com/jinzhu/gorm"
	"gorm-simple/lib"
	"time"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	DB, err = gorm.Open("mysql", "root:1513@/beego?parseTime=true&charset=utf8")

	lib.CheckErr(err)

	//DB.LogMode(true)

	var user User
	var post Post
	var role Role

	DB.DropTableIfExists(&post, &user, &role)

	DB.AutoMigrate(&role, &user, &post)

	DB.Model(new(Post)).
		AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

	DB.Model(new(User)).
		AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")
}

type TimeStamps struct {
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" sql:"index"`
}
