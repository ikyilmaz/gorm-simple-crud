package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"gorm-simple/lib"
	"gorm-simple/models"
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

	err := models.DB.
		Preload("Posts").
		Preload("Role").
		Scopes(models.DefaultUserScope).
		Find(&users).
		Error

	lib.CheckErr(err)

	for _, user := range users {
		printUser(user)
	}

	var post models.Post

	err = models.DB.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Scopes(models.DefaultUserScope)
		}).
		Find(&post, 1).
		Error

	lib.CheckErr(err)

	printPost(post)
}

func printUser(user models.User) {
	//color.Cyan("-------------USER---------------\n\n")
	//fmt.Printf("ID: %b FNAME: %s LNAME: %s EMAIL: %s\nPWDHASH: %s\n ROLE_ID: %b ROLENAME: %s",
	//	user.ID,
	//	user.FirstName,
	//	user.LastName,
	//	user.Email,
	//	user.Password,
	//	user.RoleID,
	//	user.Role.Name,
	//)
	//color.Cyan("-----------USERPOSTS------------\n\n")
	//for _, post := range user.Posts {
	//	fmt.Printf("ID: %b TITLE: %s USER_ID: %b\n", post.ID, post.Title, post.UserID)
	//}
	//color.Cyan("-------------END---------------\n\n")
}

func printPost(post models.Post) {
	//color.Cyan("-------------POST---------------\n\n")
	//fmt.Printf("ID: %b NAME: %s USER_ID: %b\n", post.ID, post.Title, post.UserID)
	//if post.User != nil {
	//	color.Cyan("-----------POSTUSER------------\n\n")
	//	fmt.Printf("USER_ID: %b USER_FNAME: %s USER_LNAME: %s EMAIL: %s\nPWDHASH: %s\n",
	//		post.User.ID,
	//		post.User.FirstName,
	//		post.User.LastName,
	//		post.User.Email,
	//		post.User.Password,
	//	)
	//}
	//color.Cyan("-------------END---------------\n\n")
}
