package models

import (
	"time"
	"github.com/jinzhu/gorm"
	"myblog/base"
	"errors"
)

type ArticleModel struct {
	Id         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Type       int       `json:"type"`
	TagId 	   int       `json:"tag_id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
func (m ArticleModel) TableName() string {
	return "article"
}

func (m *ArticleModel) GetDB() *gorm.DB {
	return base.DBMyBlog.Table(m.TableName())
}

func (m *ArticleModel) AddArticle(	) error {
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	return m.GetDB().Create(m).Error
}

func (m *ArticleModel) DeleteArticle(id int) error {
	return m.GetDB().Where("id=?", id).Delete(m).Error
}


func (m *ArticleModel) UpdateArticle(id int, data map[string]interface{}) error {
	m.UpdateTime = time.Now()
	return m.GetDB().Where("id=?", id).Updates(data).Error
}


func (m *ArticleModel) GetArticleTitle() error {
	return m.GetDB().Exec("SELECT title FROM article;").Error
}


func (m *ArticleModel) GetArticleID(id int) error {
	err := m.GetDB().Where("id=?", id).Find(m).Error
	if err != nil{
		if err == gorm.ErrRecordNotFound{
			return errors.New("文章不存在")
		}else{
			return err
		}
	}
	return nil
}


type FormatArticle struct {
	ArticleModel
	TagData TagModel `json:"tag_name"`
}

func (m *ArticleModel) GetArticlesOfTag(tag_id int) ([] FormatArticle, error) {
	articleList := make([] FormatArticle, 0)

	err :=  m.GetDB().Where("tag_id=?", tag_id).Find(&articleList).Error
	if err!=nil{
		base.Logger.Errorf("Error %s", err.Error())
		return nil, err
	}




	tagModel := new(TagModel)
	tagModel.GetDB().Where("id = ?", tag_id).Find(tagModel)
	for _, v := range articleList{
		v.TagData = *tagModel
	}



	return articleList, nil

}