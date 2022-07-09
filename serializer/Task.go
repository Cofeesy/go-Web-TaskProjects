package serializer

import "memorandumProject/model"

//Task序列器
type Task struct {
	ID        uint   `json:"id" example:"1"`
	Title     string `json:"title" example:"吃饭"`
	Content   string `json:"content" example:"睡觉"`
	View      uint64 `json:"view" example:"32"` //浏览量
	Status    int    `json:"status" example:"0"`
	CreateAt  int64  `json:"createAt"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
}

/**
 * @Author Cofeesy
 * @Description //单个task构造函数
 * @Date 23:00 2022/6/25
 * @Param item model.Task
 * @return Task
 **/
func BuildTask(item model.Task) Task {
	// 构造数据模型
	return Task{
		ID:      item.ID,
		Title:   item.Title,
		Content: item.Content,
		//View: item.View(),
		Status:    item.Status,
		StartTime: item.StartTime,
		EndTime:   item.EndTime,
	}
}

/**
 * @Author Cofessy
 * @Description //多个task构造函数
 * @Date 23:52 2022/6/25
 * @Param items []model.Task ”数据库task的数组切片“
 * @return tasks []Task ”构造器Task的数组切片“
 **/
func BuildTasks(items []model.Task) (tasks []Task) {
	for _, item := range items {
		task := BuildTask(item)
		tasks = append(tasks, task)
	}
	return tasks
}
