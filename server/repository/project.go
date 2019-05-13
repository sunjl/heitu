package repository

import (
	"github.com/jishufan/heitu/server/db"
	"time"
)

type Project struct {
	Id         int       `gorm:"column:id;primary_key" json:"id"`
	GroupName  string    `gorm:"column:group_name" json:"groupName"`
	Name       string    `gorm:"column:name" json:"name"`
	Git        string    `gorm:"column:git" json:"git"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (Project) TableName() string {
	return "projects"
}

func CreateProject(project Project) (Project, error) {
	project.CreateTime = time.Now()
	result := db.MysqlDB.Create(&project)
	return project, result.Error
}

func FindProjectById(id int) (Project, error) {
	project := Project{}
	result := db.MysqlDB.Where(&Project{Id: id}).First(&project)
	return project, result.Error
}

func FindProjectByName(name string) (Project, error) {
	project := Project{}
	result := db.MysqlDB.Where(&Project{Name: name}).First(&project)
	return project, result.Error
}

func FindAllProjectList(query map[string]interface{}, pageRequest PageRequest) ([]Project, error) {
	projects := []Project{}
	orderValue := pageRequest.Order + " " + pageRequest.Direction
	result := db.MysqlDB.Where(query).Order(orderValue).Find(&projects)
	return projects, result.Error
}

func FindProjectList(query map[string]interface{}, pageRequest PageRequest) ([]Project, error) {
	projects := []Project{}
	orderValue := pageRequest.Order + " " + pageRequest.Direction
	offsetValue := pageRequest.Page * pageRequest.Size
	limitValue := pageRequest.Size
	result := db.MysqlDB.Where(query).Order(orderValue).Offset(offsetValue).Limit(limitValue).Find(&projects)
	return projects, result.Error
}

func CountProject(query map[string]interface{}) (int, error) {
	var count int
	result := db.MysqlDB.Model(&Project{}).Where(query).Count(&count)
	return count, result.Error
}

func UpdateProject(project Project) (Project, error) {
	existingProject, err := FindProjectById(project.Id)
	if err != nil {
		return project, err
	}
	project.CreateTime = existingProject.CreateTime
	project.UpdateTime = time.Now()
	result := db.MysqlDB.Save(&project)
	return project, result.Error
}

func DeleteProjectById(id int) error {
	existingProject, err := FindProjectById(id)
	if err != nil {
		return err
	}
	result := db.MysqlDB.Delete(&existingProject)
	return result.Error
}
