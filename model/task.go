package model

import (
	// "strconv"
	"github.com/jinzhu/gorm"
)

//Task 任务模型
type Task struct {
	gorm.Model
	User      User   `gorm:"ForeignKey:Uid"` //创建该任务的用户；逻辑外键，在数据库中不显示
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"index;not null"` //index:普通索引，可以不加名字，gorm自动生成；not null:非空约束
	Status    int    //`gorm:"default:'0'"`
	Content   string `gorm:"type:longtext"`
	StartTime int64  //备忘录创建时间
	EndTime   int64  `gorm:"default:0"` //备忘录创建完成时间
}

/*
func (Task *Task) ViewAdd() uint64{
	countStr,_:=cache.RedisClient.Get(cache.TaskViewKey(Task.ID)).Result()
	count,_:strconv.ParseUint(countStr,10,64)
	return count
}
*/

/*
func (Task *Task)AddView()  {

}
*/