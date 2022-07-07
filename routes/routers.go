package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"memorandumProject/api"
	"memorandumProject/middleware"
)

/**
 * @Title NewRouter
 * @Description //路由启动
 * @Author Cofeesy 10:34 2022/7/5
 * @Param
 * @Return *gin.Engine
 **/
func NewRoouter() *gin.Engine {
	//中间件启动
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("session", store))
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		//用户注册
		v1.POST("user/register", api.UserRigister)
		//用户登录
		v1.POST("user/login", api.UserLogin)
		authored := v1.Group("/")
		authored.Use(middleware.JWT())
		{
			//创建一个任务
			authored.POST("task", api.CreateTask)
			//展示一个任务
			authored.GET("task/:id", api.ShowTask)
			//查看多个任务
			authored.GET("tasks", api.ListTasks)
			//删除一个任务
			authored.DELETE("task/:id", api.DeleteTask)
			//更新一个任务
			authored.PUT("task/:id", api.UpdateTask)
			//搜索任务
			authored.POST("search", api.SearchTasks)
		}
	}

	return r
}
