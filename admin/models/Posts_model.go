package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title, Slug, Description, Content, Picture_url string
	CategoryID                                     int
}

func CreatePost() {
	Database.Db.Create(&Post{Title: "title3", Slug: "slug3", Description: "dec3", Content: "con", Picture_url: "url", CategoryID: 54})
}
func GetAllPosts(where ...interface{}) []Post {
	var posts []Post
	Database.Db.Find(&posts, where...)
	return posts
}
func GetSinglePost(where ...interface{}) Post {
	var post Post
	Database.Db.First(&post, where...)
	return post
}
func UpdatePost(column string, value interface{}) {
	var post Post
	Database.Db.First(&post, 1)
	Database.Db.Model(&post).Update(column, value)
}
func UpdatesPost(data Post) {
	var post Post
	Database.Db.First(&post, 1)
	Database.Db.Model(&post).Updates(data)
}
func DeletePost() {
	var post Post
	Database.Db.Delete(&post, post.ID)
}
