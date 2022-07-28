package model

import (
	"PalaemonBlog/utils/errormsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCategoryStatus 查询分类是否存在 | query category status
func CheckCategoryStatus(name string) (code int) {
	var category Category
	Db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return errormsg.ErrorCategoryUsed
	}
	return errormsg.SUCCESS
}

// CreateNewCategory 添加分类 | add new category
func CreateNewCategory(data *Category) int {
	err := Db.Create(&data).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

// QuerySingleCategory 查询单个分类下的文章 | query article under category
func QuerySingleCategory(ID int) (Category, int) {
	var category Category
	Db.Where("id = ?", ID).First(&category)
	return category, errormsg.SUCCESS
}

// QueryCategoryList 查询分类列表 | query category list
func QueryCategoryList(PageSize, PageNum int) ([]Category, int64) {
	var category []Category
	var total int64
	err = Db.Limit(PageSize).Offset((PageNum - 1) * PageSize).Find(&category).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		return nil, 0
	}
	Db.Model(&category).Count(&total)
	return category, total
}

// EditCategory 编辑分类 | edit category
func EditCategory(ID int, data *Category) (code int) {
	var category Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err = Db.Model(&category).Where("id = ?", ID).Updates(maps).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

// DeleteCategory 删除分类 | delete category
func DeleteCategory(ID int) (code int) {
	var category Category
	err := Db.Where("id = ?", ID).Delete(&category).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}
