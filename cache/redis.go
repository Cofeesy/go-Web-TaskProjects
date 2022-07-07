package cache

import (
	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus"
	"memorandumProject/config"
	"strconv"
)

/**
 * @Author Cofessy
 * @Description //初始化Redis连接
 * @Date 21:40 2022/6/25
 * @Param
 * @return
 **/
func Redis() {
	db, _ := strconv.ParseUint(config.RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddr,
		Password: config.RedisPw,
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	config.RedisClient = client

}
