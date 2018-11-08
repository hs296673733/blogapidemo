package models

import (
	"github.com/jinzhu/gorm"
	"myblog/base"
)

type UserModel struct {
	Id int `json:"id"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	NickName string `json:"nick_name"`
}


func (u UserModel) TableName()string {
	return "user"
}


func (u *UserModel) GetDB() *gorm.DB {
	return base.DBMyBlog.Table(u.TableName())
}


func (u *UserModel) AddUser() error {
	return u.GetDB().Create(u).Error
}


func (u *UserModel) EditUser(id int, data map[string]interface{}) error {
	return u.GetDB().Where("id=?", id).Update(u).Error
}


func (u *UserModel) CheckUser(username, string2 string, password string) bool {
	var user UserModel
	u.GetDB().Select("id").Where(UserModel{UserName:username, PassWord:password}).First(u)
	if user.Id > 0 {
		return true
	}

	return false
}
