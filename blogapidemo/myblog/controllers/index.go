package controllers

import (
	"github.com/gin-gonic/gin"
	"myblog/base"
	"fmt"
)


type IndexReq struct {
	YouName string `json:"you_name" form:"you_name"`
}


func ApiIndex(c *gin.Context) {
	req := new(IndexReq)
	if err := c.Bind(req);err!=nil {
		base.ExceptGeneral(c, "参数错误", nil)
		base.Logger.Errorf("Params error %s", err.Error())
		return
	}

	base.Logger.Warningf("Hello world %s", req.YouName)

	base.Success(c, fmt.Sprintf("hello, world %s", req.YouName), nil)

}


