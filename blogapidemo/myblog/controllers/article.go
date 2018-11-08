package controllers

import (
	"github.com/gin-gonic/gin"
	"myblog/base"
	"fmt"
	"myblog/models"
)

// swagger:parameters AddArticleReq
type AddArticleReq struct {
	Title string `json:"title"`
	Content string `json:"content"`
	Type int `json:"type"`
}


func ApiAddArticle(c *gin.Context) {
	req := new(AddArticleReq)
	if err := c.Bind(req);err!=nil {
		base.ExceptGeneral(c, "参数错误", nil)
		base.Logger.Errorf("Params error %s", err.Error())
		return
	}

	article := new(models.ArticleModel)
	article.Content = req.Content
	article.Title = req.Title
	article.Type = req.Type
	article.AddArticle()

	base.Success(c, fmt.Sprintf("添加文章成功 %s", req.Title), article.Id)

}


type DeleteArticleReq struct {
	Id int `json:"id"`
}


func ApiDeleteArticle(c *gin.Context) {
	req := new(DeleteArticleReq)
	if err := c.Bind(req);err != nil{
		base.ExceptGeneral(c, "参数错误", nil)
		base.Logger.Errorf("Params error %s", err.Error())
		return
	}
	article := new(models.ArticleModel)
	result := article.DeleteArticle(req.Id)
	base.Success(c, fmt.Sprintf("删除文章成功"), result)
}


type UpdateArticleReq struct {
	Id int `json:"id"`
	Title string `json:"title"`
	content string `json:"content"`
}


func ApiUpdateArticle(c *gin.Context) {
	req := new(UpdateArticleReq)
	if err := c.Bind(req);err != nil{
		base.ExceptGeneral(c, "参数错误", nil)
		base.Logger.Errorf("Params error %s", err.Error())
		return
	}
	article := new(models.ArticleModel)
	article.Title = req.Title
	article.Content = req.content
	data := make(map[string]interface{})
	data["title"] = article.Title
	data["content"] = article.Content
	result := article.UpdateArticle(req.Id, data)
	base.Success(c, fmt.Sprintf("更新成功"), result)

}

type GetArticleReq struct {
	Id int `json:"id"`
	Title string `json:"title"`
	content string `json:"content"`
	Tag_id int `json:"tag_id"`
}


func ApiGetArticleTitle(c *gin.Context) {
	req := new(GetArticleReq)
	if err := c.Bind(req); err !=nil{
		base.ExceptGeneral(c, "参数错误", nil)
		base.Logger.Errorf("Params error %s", err.Error())
		return
	}
	article := new(models.ArticleModel)
	result := article.GetArticleTitle()
	base.Success(c, fmt.Sprintf("查询成功"), result)
}


func ApiGetArticle(c *gin.Context){
	req := new(GetArticleReq)
	if err := c.Bind(req); err != nil {
		base.ExceptGeneral(c, "参数错误", nil)
		base.Logger.Errorf("Params error %s", err.Error())
		return
	}
	article := new(models.ArticleModel)
	article.Title = req.Title
	article.Content = req.content
	result := article.GetArticleID(req.Id)
	base.Success(c, fmt.Sprintf("查询成功"), result)
}


type GetArticleOfTagReq struct {
	TagId int `json:"tag_id"`
}

func ApiGetArticleOfTag(c *gin.Context) {
	req := new(GetArticleOfTagReq)
	if err := c.Bind(req); err != nil {
		base.ExceptGeneral(c, "参数错误", nil)
		base.Logger.Errorf("Params error %s", err.Error())
		return
	}

	article := new(models.ArticleModel)
	result , err := article.GetArticlesOfTag(req.TagId)
	if err != nil{
		base.ExceptLogic(c, "查询错误", err)
		return
	}
	base.Success(c, fmt.Sprintf("查询成功"), result)

}