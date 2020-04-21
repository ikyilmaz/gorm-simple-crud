package models

import (
	"github.com/jinzhu/gorm"
	"gorm-simple-crud/lib"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	DB, err = gorm.Open("mysql", "root:1513@/beego?parseTime=true&charset=utf8")

	lib.CheckErr(err)

	DB.LogMode(true)

	DB.DropTableIfExists(new(Post), new(User))

	DB.AutoMigrate(new(User), new(Post))

	DB.Model(new(Post)).
		AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}
