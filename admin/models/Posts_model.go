package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title, Slug, Description, Content, Picture_url string
	CategoryID                                     int
}

func (post Post) CreatePost() {
	Database.Db.Create(&post)
}
func (post Post) GetAllPosts(where ...interface{}) []Post {
	var posts []Post
	Database.Db.Find(&posts, where...)
	return posts
}
func (post Post) GetSinglePost(where ...interface{}) Post {
	Database.Db.First(&post, where...)
	return post
}
func UpdatePost(column string, value interface{}) {
	var post Post
	Database.Db.First(&post, 1)
	Database.Db.Model(&post).Update(column, value)
}
func (post Post) UpdatesPost(data Post) {
	Database.Db.First(&post, 1)
	Database.Db.Model(&post).Updates(data)
}
func (post Post) DeletePost() {
	Database.Db.Delete(&post, post.ID)
}
