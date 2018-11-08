package controllers

import (
	"github.com/gin-gonic/gin"
	"myblog/base"
	"myblog/models"
)

type GetTagReq struct {
	Id int `json:"id"`
	Tag_name string `json:"tag_name"`
}


func ApiGetTags(c *gin.Context) {
	req := new(GetTagReq)
	if err := c.Bind(req); err != nil {
		base.ExceptGeneral(c, "参数错误", nil)
		base.Logger.Errorf("Params error %s", err.Error())
		return
	}
	tag := new(models.TagModel)
	tag.GetTags()
}


