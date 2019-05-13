package repository

import (
	"github.com/jishufan/heitu/server/db"
	"time"
)

type Task struct {
	Id         int       `gorm:"column:id;primary_key" json:"id"`
	HostName   string    `gorm:"column:host_name" json:"hostName"`
	IP         string    `gorm:"column:ip" json:"ip"`
	Name       string    `gorm:"column:name" json:"name"`
	Content    string    `gorm:"column:content" json:"content"`
	Status     string    `gorm:"column:status" json:"status"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (Task) TableName() string {
	return "tasks"
}

func CreateTask(task Task) (Task, error) {
	task.CreateTime = time.Now()
	result := db.MysqlDB.Create(&task)
	return task, result.Error
}

func FindTaskById(id int) (Task, error) {
	task := Task{}
	result := db.MysqlDB.Where(&Task{Id: id}).First(&task)
	return task, result.Error
}

func FindAllTaskList(query map[string]interface{}, pageRequest PageRequest) ([]Task, error) {
	tasks := []Task{}
	orderValue := pageRequest.Order + " " + pageRequest.Direction
	result := db.MysqlDB.Where(query).Order(orderValue).Find(&tasks)
	return tasks, result.Error
}

func FindTaskList(query map[string]interface{}, pageRequest PageRequest) ([]Task, error) {
	tasks := []Task{}
	orderValue := pageRequest.Order + " " + pageRequest.Direction
	offsetValue := pageRequest.Page * pageRequest.Size
	limitValue := pageRequest.Size
	result := db.MysqlDB.Where(query).Order(orderValue).Offset(offsetValue).Limit(limitValue).Find(&tasks)
	return tasks, result.Error
}

func CountTask(query map[string]interface{}) (int, error) {
	var count int
	result := db.MysqlDB.Model(&Task{}).Where(query).Count(&count)
	return count, result.Error
}

func UpdateTask(task Task) (Task, error) {
	existingTask, err := FindTaskById(task.Id)
	if err != nil {
		return task, err
	}
	task.CreateTime = existingTask.CreateTime
	task.UpdateTime = time.Now()
	result := db.MysqlDB.Save(&task)
	return task, result.Error
}

func DeleteTaskById(id int) error {
	existingTask, err := FindTaskById(id)
	if err != nil {
		return err
	}
	result := db.MysqlDB.Delete(&existingTask)
	return result.Error
}
