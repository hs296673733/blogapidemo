package controllers

import (
	"github.com/gin-gonic/gin"
	"myblog/base"
	"myblog/models"
)

type RegisterUserReq struct {
	Id int `json:"id"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	NickName string `json:"nick_name"`
}

func ApiRegisterUser(c *gin.Context) {
	req := new(RegisterUserReq)
	if err := c.Bind(req); err != nil {
		base.ExceptGeneral(c, "参数错误", nil)
		base.Logger.Errorf("Params error %s", err.Error())
		return
	}

	user := new(models.UserModel)
	user.UserName = req.UserName
	user.PassWord = req.PassWord
	user.NickName = req.NickName
	user.AddUser()
}


type EditUserReq struct {
	Id int `json:"id"`
	Password string `json:"pass_word"`
	NickName string `json:"nick_name"`
}

func ApiEdituser(c *gin.Context) {
	req := new(EditUserReq)
	if err := c.Bind(req); err != nil {
		base.ExceptGeneral(c, "参数错误", nil)
		base.Logger.Errorf("Params error %s", err.Error())
		return
	}

	user := new(models.UserModel)
	user.PassWord = req.Password
	user.NickName = req.NickName
	data := make(map[string]interface{})
	data["password"] = user.PassWord
	data["nickname"] = user.NickName
	user.EditUser(req.Id, data)
}