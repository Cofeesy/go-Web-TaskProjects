package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//数据库User表映射
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PassWordDigest string
}

//密码加密难度
const (
	PassWordCost = 20
)

/**
 * @Author Cofessy
 * @Description //对传进来的将要存入数据库的用户密码加密
 * @Date 23:13 2022/6/24
 * @Param password string
 * @return error
 **/
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PassWordDigest = string(bytes)
	return nil
}

/**
 * @Author Cofeesy
 * @Description //对存入数据库的用户密码进行解密
 * @Date 23:15 2022/6/24
 * @Param password string
 * @return bool
 **/
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWordDigest), []byte(password))
	return err == nil
}
