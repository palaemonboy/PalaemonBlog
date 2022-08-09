package model

import (
	"PalaemonBlog/utils/errormsg"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200);not null" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(1000)" json:"img"`
}

// CreateNewArticle 添加文章 | add new article
func CreateNewArticle(data *Article) int {
	err := Db.Create(&data).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

// QueryArticleList 查询文章列表 | query article list
func QueryArticleList(PageSize, PageNum int) ([]Article, int, int64) {
	var article []Article
	var total int64
	err = Db.Preload("Category").Limit(PageSize).Offset((PageNum - 1) * PageSize).Find(&article).Error
	if err != gorm.ErrRecordNotFound && err != nil {
		return nil, errormsg.ERROR, 0
	}
	Db.Model(&article).Count(&total)
	return article, errormsg.SUCCESS, total
}

// QuerySingleArticle 查询单个文章 | query single article
func QuerySingleArticle(ID int) (Article, int) {
	var article Article
	err := Db.Preload("Category").Where("id = ?", ID).First(&article).Error
	if err != nil {
		return article, errormsg.ErrorArticleNotExist
	}
	return article, errormsg.SUCCESS
}

// QueryArticlesByCategory 查询分类下的所有文章 | query all articles under category
func QueryArticlesByCategory(ID, PageSize, PageNum int) ([]Article, int, int64) {
	var article []Article
	var total int64
	err := Db.Preload("Category").Where("cid = ?", ID).Limit(PageSize).Offset((PageNum - 1) * PageSize).Find(&article).Error
	Db.Model(&article).Where("cid = ?", ID).Count(&total)
	if err != nil {
		return nil, errormsg.ErrorCategoryNotExist, 0
	}
	return article, errormsg.SUCCESS, total
}

// EditArticle 编辑文章 | edit article
func EditArticle(ID int, data *Article) (code int) {
	var article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err = Db.Model(&article).Where("id = ?", ID).Updates(maps).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

// DeleteArticle 删除文章 | delete article
func DeleteArticle(ID int) (code int) {
	var article Article
	err := Db.Where("id = ?", ID).Delete(&article).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}
