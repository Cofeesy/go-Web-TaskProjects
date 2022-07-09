package service

import (
	"github.com/jinzhu/gorm"
	"memorandumProject/model"
	"memorandumProject/pkg/e"
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
	code := e.SUCCESS
	model.DB.Model(&model.User{}).Where("user_name=?", sevice.UserName).First(&user).Count(&count)
	if count == 1 {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user.UserName = sevice.UserName

	//加密密码
	if err := user.SetPassword(sevice.PassWord); err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (service *UserService) Login() serializer.Response {
	var user model.User
	code := e.SUCCESS
	//数据库查询是否存在，不存在则数据库报错
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		//gorm函数判断错误类型是否是该用户载数据库没有记录
		if gorm.IsRecordNotFoundError(err) {
			//错误日志打印到logs文件
			utils.LogrusObj.Info(err)
			//规范错误码
			code = e.ErrorNotExistUser
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		//用户存在，依然有错误
		//记录日志
		utils.LogrusObj.Info(err)
		//规范错误码
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//如果都没错误，则进行密码验证，密码验证成功,则进行token颁发
	if !user.CheckPassword(service.PassWord) {
		//规范错误码
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//token颁发
	token, err := utils.GenToken(user.ID, service.UserName, 1)
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
		Msg: e.GetMsg(code),
	}

}
