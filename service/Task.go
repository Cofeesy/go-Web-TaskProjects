/**
  @author:Cofessy
  @data:2022/6/24
  @note:service通过gorm与数据库交互
**/
package service

import (
	"memorandumProject/model"
	"memorandumProject/pkg/e"
	"memorandumProject/pkg/utils"
	"memorandumProject/serializer"
	"time"
)

//创建任务的服务
type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"`
}

//展示任务详情的服务
type ShowTaskService struct {
}

//删除任务的服务
type DeleteTaskService struct {
}

//更新任务的服务
type UpdateTaskService struct {
	ID      uint   `json:"ID" form:"id"`
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"`
}

//分页返回的服务
type ListTasksService struct {
	Limit int `json:"limit" form:"limit"`
	Start int `json:"start" form:"limit"`
}

type SearchTaskService struct {
	Info string `json:"info" form:"info"`
}

/**
 * @Title CreateTask
 * @Description //service层创建任务-->对数据库进行新增
 * @Author Cofeesy 12:29 2022/6/26
 * @Param id uint
 * @Return serializer.Response
 **/
func (service *CreateTaskService) Create(id uint) serializer.Response {
	var user model.User
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Status:    service.Status,
		Content:   service.Content,
		StartTime: time.Now().Unix(),
	}
	code := e.SUCCESS
	err := model.DB.Create(&task).Error
	//err不为空，返回错误
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    "创建备忘录失败",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    e.GetMsg(code),
	}
}

/**
 * @Title Delete
 * @Description //service层删除对应的记录
 * @Author Cofeesy 15:54 2022/6/27
 * @Param id string
 * @Return serializer.Responsea
 **/
func (sevice *DeleteTaskService) Delete(id string) serializer.Response {
	var task model.Task
	code := e.SUCCESS
	//删除之前先查找对应id(主键唯一性)的记录是否存在
	err := model.DB.First(&task, id).Error
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//找到则进行删除
	err = model.DB.Delete(&task).Error
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//删除成功返回
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

/**
 * @Title Update
 * @Description //service层修改task
 * @Author Cofeesy 16:02 2022/6/27
 * @Param service *UpdateTaskService
 * @Return serializer.Response
 **/
func (service *UpdateTaskService) Update(id string) serializer.Response {
	var task model.Task
	model.DB.Model(model.Task{}).Where("id=?", id).First(&task)
	task.Content = service.Content
	task.Status = service.Status
	task.Title = service.Title
	code := e.SUCCESS
	err := model.DB.Save(&task).Error
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   "修改成功",
	}

}

/**
 * @Title search
 * @Description //TODO
 * @Author Cofeesy 18:40 2022/6/27
 * @Param Uid uint
 * @Return serializer.Response
 **/
func (service *SearchTaskService) Search(Uid uint) serializer.Response {
	var tasks []model.Task
	code := e.SUCCESS
	model.DB.Where("Uid=?", Uid).Preload("User").First(&tasks)
	err := model.DB.Model(&model.Task{}).Where("title LIKE ? OR content LIKE?",
		"%"+service.Info+"%"+service.Info+"%").Find(&tasks).Error
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildTasks(tasks),
	}
}

/**
 * @Title Show
 * @Description //处理api层通过查看任务-->
 * @Author Cofeesy 19:23 2022/6/27
 * @Param id string
 * @Return serializer.Response
 **/
func (service *ShowTaskService) Show(id string) serializer.Response {
	var task model.Task
	code := e.SUCCESS
	//数据库查找
	err := model.DB.First(&task, id).Error
	if err != nil {
		utils.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//增加点击量
	//task.Addview()
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
		Msg:    e.GetMsg(code),
	}
}

/**
 * @Title List
 * @Description //TODO
 * @Author Cofeesy 0:04 2022/6/27
 * @Param id int
 * @Return serializer.Response
 **/
func (service *ListTasksService) List(id uint) serializer.Response {
	var tasks []model.Task
	var total int64
	//
	if service.Limit == 0 {
		service.Limit = 15
	}
	model.DB.Model(model.Task{}).Preload("User").Where("uid=?", id).Count(&total).
		Limit(service.Limit).Offset((service.Start - 1) * service.Limit).Find(&tasks)
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(total))
}
