package model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

//全局数据库操作对象DB
var DB *gorm.DB

/**
 * @Author Cofeesy
 * @Description // gormv2数据库连接以及日志打印
 * @Date 18:12 2022/6/25
 * @Param conn string
 * @return
 **/
func Database(conn string) {
	//全局日志输出
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	//gormv2数据库连接以及配置
	db, err := gorm.Open(mysql.New(mysql.Config{
		//DSN data source name
		DSN: conn,
		//string类型字段的默认长度
		DefaultStringSize: 256,
		//禁用datetime精度
		DisableDatetimePrecision: true,
		//重命名索引时采用删除并新建的方式
		DontSupportRenameIndex:  true,
		DontSupportRenameColumn: true,
		//根据版本自动配置
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		//打印日志
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //表名单数，不加“s”
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	//设置连接池
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 60)
	DB = db
	migration()
}

/*
gorm:v1
func Database(conn string) {
	//初级排bug--有错误就逐步向上发现错误用打印错误方式来实现
	//fmt.Println("conn", conn)
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic("mysql连接错误")
	}
	fmt.Println("数据库连接成功")

	db.LogMode(true)
	if gin.Mode() == "release" {
		db.LogMode(false)
	}

	//表名单数
	db.SingularTable(true)
	//设置连接池
	db.DB().SetMaxIdleConns(20)
	//设置最大连接数
	db.DB().SetMaxOpenConns(100)
	//连接时间
	db.DB().SetConnMaxLifetime(time.Second * 30)
	//将db赋值给全局变量DB
	DB = db

	migration()
}
*/
