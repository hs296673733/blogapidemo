package main

import (
	"myblog/conf"
	"myblog/base"
	"myblog/controllers"
)

func main(){
	conf.InitFileConfig("./config.json")
	base.InitLogger()
	base.InitDB()
	base.InitGinServer()
	controllers.SetupRouter(base.GinServer)
	base.GinServer.Run(conf.Config.ListenAddr)
}
