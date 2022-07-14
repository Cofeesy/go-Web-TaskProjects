package config

import (
	"gopkg.in/ini.v1"
	"memorandumProject/model"
	"memorandumProject/pkg/utils"
	"strings"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

/**
 * @Author Cofeesy
 * @Description //读取配置文件并连接数据库
 * @Date 18:10 2022/6/23
 * @Param
 * @return
 **/
func Init() {
	//加载并读取配置文件内容
	file, err := ini.Load("./config/conf.ini")
	if err != nil {
		utils.LogrusObj.Info(err)
		panic(err)
	}

	LoadServer(file)
	LoadMysql(file)

	dsn := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	model.Database(dsn)

}

/**
 * @Author Cofeesy
 * @Description // 加载server文件的配置
 * @Date 18:12 2022/6/23
 * @Param file *ini.File
 * @return
 **/
func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

/**
 * @Author Cofessy
 * @Description //加载mysql文件配置
 * @Date 22:40 2022/6/24
 * @Param file *ini.File
 * @return
 **/
func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
