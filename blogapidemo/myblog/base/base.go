package base

import "github.com/gin-gonic/gin"
import "github.com/gin-contrib/cors"

var GinServer *gin.Engine


func InitGinServer(){
	GinServer = gin.Default()

	GinServer.Use(cors.Default())
	//GinServer.SetMode(gin.ReleaseMode)


}