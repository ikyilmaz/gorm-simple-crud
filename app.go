package main

import (
	"fmt"
	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"gorm-simple-crud/models"
)

func main() {
	defer println("done!")
	defer models.DB.Close()

	models.DB.Create(&models.User{
		FirstName: "john",
		LastName:  "doe",
		Email:     "ismail@example.com",
		Password:  "87654321",
	})

	models.DB.Create(&models.User{
		FirstName: "becka",
		LastName:  "grimes",
		Email:     "becka@example.com",
		Password:  "87654321",
	})

	models.DB.Create(&models.Post{
		Title:  "Something",
		UserId: 1,
	})

	models.DB.Create(&models.Post{
		Title:  "Another thing",
		UserId: 1,
	})

	models.DB.Create(&models.Post{
		Title:  "Good morning",
		UserId: 2,
	})

	models.DB.Create(&models.Post{
		Title:  "Good afternoon",
		UserId: 2,
	})

	var users []models.User

	models.DB.Preload("Posts").Scopes(models.DefaultUserScope).Find(&users)

	for _, user := range users {
		printUser(user)
	}

	var post models.Post

	models.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Scopes(models.DefaultUserScope)
	}).Find(&post, 1)

	printPost(post)
}

func printUser(user models.User) {
	color.Cyan("-------------USER---------------\n\n")
	fmt.Printf("ID: %b FNAME: %s LNAME: %s EMAIL: %s\nPWDHASH: %s\n",
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
	)
	color.Cyan("-----------USERPOSTS------------\n\n")
	for _, post := range user.Posts {

		fmt.Printf("ID: %b TITLE: %s USER_ID: %b\n", post.ID, post.Title, post.UserId)
	}
	color.Cyan("-------------END---------------\n\n")
}

func printPost(post models.Post) {
	color.Cyan("-------------POST---------------\n\n")
	fmt.Printf("ID: %b NAME: %s USER_ID: %b\n", post.ID, post.Title, post.UserId)
	if post.User != nil {
		color.Cyan("-----------POSTUSER------------\n\n")
		fmt.Printf("USER_ID: %b USER_FNAME: %s USER_LNAME: %s EMAIL: %s\nPWDHASH: %s\n",
			post.User.ID,
			post.User.FirstName,
			post.User.LastName,
			post.User.Email,
			post.User.Password,
		)
	}
	color.Cyan("-------------END---------------\n\n")
}
