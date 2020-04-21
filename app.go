package main

import (
	"fmt"
	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"gorm-simple/lib"
	"gorm-simple/models"
	"gorm-simple/models/scopes"
)

func main() {
	defer println("done!")
	defer models.DB.Close()

	models.DB.Create(&models.Role{
		Name:        "user",
		Description: "readonly",
	})

	models.DB.Create(&models.Role{
		Name:        "author",
		Description: "read write own posts",
	})

	models.DB.Create(&models.Role{
		Name:        "admin",
		Description: "read write",
	})

	models.DB.Create(&models.User{
		FirstName: "john",
		LastName:  "doe",
		Email:     "ismail@example.com",
		Password:  "87654321",
		RoleID:    2,
	})

	models.DB.Create(&models.User{
		FirstName: "becka",
		LastName:  "grimes",
		Email:     "becka@example.com",
		Password:  "87654321",
		RoleID:    2,
	})

	models.DB.Create(&models.Post{
		Title:  "Something",
		UserID: 1,
	})

	models.DB.Create(&models.Post{
		Title:  "Another thing",
		UserID: 1,
	})

	models.DB.Create(&models.Post{
		Title:  "Good morning",
		UserID: 2,
	})

	models.DB.Create(&models.Post{
		Title:  "Good afternoon",
		UserID: 2,
	})

	var users []models.User

	models.DB.LogMode(true)
	err := models.DB.
		Preload("Posts").
		Preload("Role").
		Scopes(scopes.PublicUserScope, scopes.DefaultUserScope).
		Find(&users).
		Error

	lib.CheckErr(err)

	models.DB.LogMode(false)

	fmt.Printf("%v\n", users[0].Role)

	for _, user := range users {
		printUser(user)
	}

	var post models.Post

	err = models.DB.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Scopes(scopes.PublicUserScope, scopes.DefaultUserScope)
		}).
		Find(&post, 1).
		Error

	lib.CheckErr(err)

	printPost(post)
}

func printUser(user models.User) {
	color.Cyan("-------------USER---------------\n\n")
	fmt.Printf("ID: %d FNAME: %s LNAME: %s EMAIL: %s\nPWDHASH: %s\nROLE_ID: %d ROLENAME: %s\n",
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.RoleID,
		user.Role.Name,
	)
	color.Cyan("-----------USERPOSTS------------\n\n")
	for _, post := range user.Posts {
		fmt.Printf("ID: %d TITLE: %s USER_ID: %d\n", post.ID, post.Title, post.UserID)
	}
	color.Cyan("-------------END---------------\n\n")
}

func printPost(post models.Post) {
	//color.Cyan("-------------POST---------------\n\n")
	//fmt.Printf("ID: %d NAME: %s USER_ID: %d\n", post.ID, post.Title, post.UserID)
	//if post.User != nil {
	//	color.Cyan("-----------POSTUSER------------\n\n")
	//	fmt.Printf("USER_ID: %d USER_FNAME: %s USER_LNAME: %s EMAIL: %s\nPWDHASH: %s\n",
	//		post.User.ID,
	//		post.User.FirstName,
	//		post.User.LastName,
	//		post.User.Email,
	//		post.User.Password,
	//	)
	//}
	//color.Cyan("-------------END---------------\n\n")
}
