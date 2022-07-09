/**
  @author:Cofeesy
  @data:2022/6/24
  @note:api接口
**/
package api

import (
	"github.com/gin-gonic/gin"
	"memorandumProject/pkg/utils"
	"memorandumProject/service"
)

/**
 * @Title CreateTask
 * @Description //处理创建任务的请求路由函数，调用sevice层的创建逻辑函数，并打包返回对应json数据
 * @Author Cofeesy 12:27 2022/6/26
 * @Param c *gin.Context
 * @Return
 **/
func CreateTask(c *gin.Context) {
	//创建一个service层的操作对象
	creatTask := service.CreateTaskService{}
	//从请求头部得到Authorization的值(用户登录创建的token值)传入进行jwt验证，该方法返回自定义的claim结构体和错误类型
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))

	if err := c.ShouldBind(&creatTask); err != nil {
		utils.LogrusObj.Info(err)
		//返回的错误是参数校验的错误，需要对这个错误进行翻译
		c.JSON(400, ErrorResponse(err))
	} else {
		res := creatTask.Create(claim.Id)
		c.JSON(200, res)
	}
}

/**
 * @Title DeleteTask
 * @Description //处理删除任务的请求路由函数，调用sevice层的删除逻辑函数，并打包返回对应json数据
 * @Author Cofeesy 18:44 2022/6/27
 * @Param c *gin.Context
 * @Return
 **/
func DeleteTask(c *gin.Context) {
	DeleteTaskService := service.DeleteTaskService{}
	res := DeleteTaskService.Delete(c.Param("id"))
	c.JSON(200, res)
}

/**
 * @Title UpdateTask
 * @Description //处理更新任务的请求路由函数，调用sevice层的更新逻辑函数，并打包返回对应json数据
 * @Author Cofeesy 18:47 2022/6/27
 * @Param c *gin.Context
 * @Return
 **/
func UpdateTask(c *gin.Context) {
	UpdateTaskService := service.UpdateTaskService{}
	//只有创作该task的人才有权限更新修改该task
	_, _ = utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&UpdateTaskService); err != nil {
		c.JSON(400, ErrorResponse(err))
		utils.LogrusObj.Info(err)
	} else {
		//Param传过来的id是string类型
		res := UpdateTaskService.Update(c.Param("id"))
		c.JSON(200, res)
	}
}

/**
 * @Title SearchTasks
 * @Description //处理查找任务请求路由函数，调用sevice层的查找逻辑函数，并打包返回对应json数据
 * @Author Cofeesy 23:28 2022/7/4
 * @Param c *gin.Context
 * @Return
 **/
func SearchTasks(c *gin.Context) {
	searchTaskService := service.SearchTaskService{}
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchTaskService); err != nil {
		c.JSON(400, ErrorResponse(err))
		utils.LogrusObj.Info(err)
	} else {
		res := searchTaskService.Search(chaim.Id)
		c.JSON(200, res)
	}
}

/**
 * @Title ShowTask
 * @Description //处理展示单个任务的请求路由函数，调用service层的分页逻辑函数
 * @Author Cofeesy 19:00 2022/6/27
 * @Param c *gin.Context
 * @Return
 **/
func ShowTask(c *gin.Context) {
	ShowTaskService := service.ShowTaskService{}
	res := ShowTaskService.Show(c.Param("id"))
	//api层返回正确
	c.JSON(200, res)
}

/**
 * @Title  ListTasks
 * @Description //处理分页的请求路由函数，调用service层的分页逻辑函数
 * @Author Cofeesy 11:36 2022/7/5
 * @Param c *gin.Context
 * @Return
 **/
func ListTasks(c *gin.Context) {
	listService := service.ListTasksService{}
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listService); err != nil {
		c.JSON(400, ErrorResponse(err))
		utils.LogrusObj.Info(err)
	} else {
		res := listService.List(chaim.Id)
		c.JSON(200, res)
	}
}
