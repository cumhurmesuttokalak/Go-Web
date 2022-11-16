package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title, Slug string
}

func (category Category) CreateCategory() {
	Database.Db.Create(&category)
}
func (category Category) GetAllCategories(where ...interface{}) []Category {
	var categories []Category
	Database.Db.Find(&categories, where...)
	return categories
}
func (category Category) GetSingleCategory(where ...interface{}) Category {
	Database.Db.First(&category, where...)
	return category
}
func (category Category) UpdateCategory(column string, value interface{}) {
	Database.Db.First(&category, 1)
	Database.Db.Model(&category).Update(column, value)
}
func (category Category) UpdatesCategory(data Category) {
	Database.Db.First(&category, 1)
	Database.Db.Model(&category).Updates(data)
}
func (category Category) DeleteCategory() {
	Database.Db.Delete(&category, category.ID)
}
