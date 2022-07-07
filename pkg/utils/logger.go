package utils

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"time"
)

var LogrusObj *logrus.Logger

/**
 * @Title init
 * @Description //日志初始化
 * @Author Cofeesy 19:05 2022/6/30
 * @Param
 * @Return
 **/
func init() {
	if LogrusObj != nil {
		src, _ := setOutPutFile()
		//设置输出
		LogrusObj.Out = src
		return
	}

	//实例化一个logger
	logger := logrus.New()
	src, _ := setOutPutFile()
	//设置输出
	logger.Out = src
	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	LogrusObj = logger
}

/**
 * @Title setOutPutFile()
 * @Description //日志输出
 * @Author Cofeesy 14:33 2022/7/6
 * @Param
 * @Return *os.File, error
 **/
func setOutPutFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	//需要创建的日志文件的拼接
	logFileName := now.Format("2006-01-02") + ".log"
	//日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		//创建日志文件
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return src, nil
}
