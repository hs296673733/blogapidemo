package models

import (
	"github.com/jinzhu/gorm"
	"myblog/base"
)

type TagModel struct {
	Id int `json:"id"`
	TagName string `json:"tag_name"`
}


func (t TagModel) TableName() string {
	return "tag"
}


func(t *TagModel) GetDB() *gorm.DB {
	return base.DBMyBlog.Table(t.TableName())
}


func (t *TagModel) GetTags() error {
	return t.GetDB().Find(t).Error
}

