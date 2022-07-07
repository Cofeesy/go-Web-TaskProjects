package service

import (
	"github.com/jinzhu/gorm"
	"memorandumProject/model"
	"memorandumProject/pkg/utils"
	"memorandumProject/serializer"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15"`
	PassWord string `form:"password" json:"password" binding:"required,min=5,max=16"`
}

//用户注册
func (sevice *UserService) Register() serializer.Response {
	var user model.User
	var count int64
	model.DB.Model(&model.User{}).Where("user_name=?", sevice.UserName).First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "已经存在该用户",
		}
	}
	user.UserName = sevice.UserName

	//加密密码
	if err := user.SetPassword(sevice.PassWord); err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    err.Error(),
		}
	}
	//创建用户
	if err := model.DB.Create(user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    err.Error(),
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "用户注册成功",
	}
}

func (service *UserService) Login() serializer.Response {
	var user model.User
	//数据库查询是否存在，不存在则数据库报错
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status: 400,
				Msg:    "该用户不存在，请先登录",
			}
		}
		//用户存在，依然有错误
		return serializer.Response{
			Status: 500,
			Msg:    "其他错误",
		}
	}
	//如果都没错误，则进行密码验证，密码验证成功,则进行token颁发
	if user.CheckPassword(service.PassWord) == false {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}
	//token颁发
	token, err := utils.GenToken(user.ID, service.UserName)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "Token颁发错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
		Msg: "登录成功",
	}

}
