package controllers

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine)  {

	router.GET("/", ApiIndex)

	articleGroup := router.Group("/article")

	{
		articleGroup.POST("/add", ApiAddArticle)
		articleGroup.POST("/delete", ApiDeleteArticle)
		articleGroup.POST("/update", ApiUpdateArticle)
		articleGroup.POST("/get_list", ApiGetArticleTitle)
		articleGroup.POST("/get_detail", ApiGetArticle)
		articleGroup.POST("/get_article_by_tag", ApiGetArticleOfTag)

	}

	tagGroup := router.Group("tag")

	{
		tagGroup.POST("get", ApiGetTags)
	}
}