package main

import (
	"memorandumProject/config"
	"memorandumProject/routes"
)

func main() { //swagger：http://localhost:3000/swagger/index.html
	//读取并初始化配置文件
	config.Init()
	//路由启动 转载路由：swag init -g main.go 需要注释
	r := routes.NewRoouter()
	//读取配置的端口
	r.Run(config.HttpPort)
}
