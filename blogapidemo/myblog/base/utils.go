package base

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
    Status    ApiStatus      `json:"status"`
    Message   string		 `json:"message"`
    Data      interface{}		 `json:"data"`
    Error     string 			 `json:"error"`
}

type ApiStatus int

const (
	ApiSuccess = ApiStatus(0)
	ApiGenericError = ApiStatus(500)
	ApiLogicError = ApiStatus(5001)
)

func Success(ctx *gin.Context, message string, data interface{}){
	rsp := &Response{
		ApiSuccess,
		message,
		data,
		"",
	}
	ctx.JSON(http.StatusOK, rsp)
}

func ExceptGeneral(ctx *gin.Context, message string, err error){
	var errStr string
	if err!=nil{
		errStr = err.Error()
	}

	rsp := Response{
		Status:  ApiGenericError,
		Message: message,
		Error:  errStr,
	}
	ctx.JSON(http.StatusOK, rsp)
}

func ExceptLogic(ctx *gin.Context, message string,  err error){
	var errStr string
	if err!=nil{
		errStr = err.Error()
	}


	rsp := Response{
		Status:  ApiLogicError,
		Message: message,
		Error:   errStr,
	}

	ctx.JSON(http.StatusOK, rsp)
}