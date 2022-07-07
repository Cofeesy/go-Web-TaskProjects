package api

import (
	"github.com/gin-gonic/gin"
	"memorandumProject/service"
)

/**
 * @Title UserRigister
 * @Description //api层用户注册绑定参数，响应错误
 * @Author Cofeesy 23:17 2022/7/4
 * @Param c *gin.Context
 * @Return
 **/

func UserRigister(c *gin.Context) {
	//创建一个UserService实例,调用service层的Register()方法
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister); err != nil {
		res := userRegister.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

/**
 * @Title UserLogin
 * @Description //api层用户登录绑定参数，响应错误
 * @Author Cofeesy 23:18 2022/7/4
 * @Param c *gin.Context
 * @Return
 **/

func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err != nil {
		res := userLogin.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}
