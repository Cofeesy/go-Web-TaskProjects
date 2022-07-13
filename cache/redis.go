package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
	"memorandumProject/pkg/utils"
	"strconv"
)

//RedisClient Redis缓存客户端单例
var (
	RedisClient *redis.Client
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func init() {
	file, err := ini.Load("./config/conf.ini")
	if err != nil {
		utils.LogrusObj.Info(err)
	}
	LoadRedis(file)
	ConnRedis()
}

/**
 * @Author Cofessy
 * @Description //初始化Redis连接
 * @Date 21:40 2022/6/25
 * @Param
 * @return
 **/
func ConnRedis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	//创建redis实例
	client := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPw,
		DB:       int(db),
	})
	res, err := client.Ping().Result()
	if err != nil {
		utils.LogrusObj.Info(err)
		panic(err)
	}
	fmt.Println(res)
	RedisClient = client
}

/**
 * @Author Cofeesy
 * @Description //加载Redis配置文件
 * @Date 21:42 2022/6/25
 * @Param file *ini.File
 * @return
 **/
func LoadRedis(file *ini.File) {
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
}
