package repository

import (
	"github.com/jishufan/heitu/server/db"
	"time"
)

type Config struct {
	Id          int       `gorm:"column:id;primary_key" json:"id"`
	ProjectId   int       `gorm:"column:project_id" json:"projectId"`
	ProjectName string    `gorm:"column:project_name" json:"projectName"`
	Version     string    `gorm:"column:version" json:"version"`
	Environment string    `gorm:"column:environment" json:"environment"`
	FileName    string    `gorm:"column:file_name" json:"fileName"`
	Content     string    `gorm:"column:content" json:"content"`
	CreateTime  time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (Config) TableName() string {
	return "configs"
}

var (
	CONFIG_BASIC_COLUMNS = []string{
		"configs.id", "configs.project_id as project_id", "projects.name as project_name",
		"configs.version", "configs.environment", "configs.file_name",
		"configs.create_time", "configs.update_time",
	}
	CONFIG_FULL_COLUMNS = []string{
		"configs.id", "configs.project_id as project_id", "projects.name as project_name",
		"configs.version", "configs.environment", "configs.file_name", "configs.content",
		"configs.create_time", "configs.update_time",
	}
)

func mapConfigQuery(query map[string]interface{}) map[string]interface{} {
	mappedQuery := make(map[string]interface{})
	for key, value := range query {
		var mappedKey string
		if key == "project_name" {
			mappedKey = "projects.name"
		} else {
			mappedKey = "configs." + key
		}
		mappedQuery[mappedKey] = value
	}
	return mappedQuery
}

func mapConfigPageRequest(pageRequest PageRequest) PageRequest {
	order := pageRequest.Order
	var mappedOrder string
	if order == "project_name" {
		mappedOrder = "projects.name"
	} else {
		mappedOrder = "configs." + order
	}
	return PageRequest{
		Order:     mappedOrder,
		Direction: pageRequest.Direction,
		Page:      pageRequest.Page,
		Size:      pageRequest.Size,
	}
}

func CreateConfig(config Config) (Config, error) {
	config.CreateTime = time.Now()
	result := db.MysqlDB.Create(&config)
	return config, result.Error
}

func FindConfigById(id int) (Config, error) {
	config := Config{}
	result := db.MysqlDB.Table("configs").Select(CONFIG_FULL_COLUMNS).
		Joins("left join projects on configs.project_id = projects.id").
		Where("configs.id = ?", id).First(&config)
	return config, result.Error
}

func FindConfig(query map[string]interface{}) (Config, error) {
	config := Config{}
	mappedQuery := mapConfigQuery(query)
	result := db.MysqlDB.Table("configs").Select(CONFIG_FULL_COLUMNS).
		Joins("left join projects on configs.project_id = projects.id").
		Where(mappedQuery).First(&config)
	return config, result.Error
}

func FindAllConfigList(query map[string]interface{}, pageRequest PageRequest) ([]Config, error) {
	configs := []Config{}
	mappedQuery := mapConfigQuery(query)
	mappedPageRequest := mapConfigPageRequest(pageRequest)
	mappedOrderValue := mappedPageRequest.Order + " " + mappedPageRequest.Direction
	result := db.MysqlDB.Table("configs").Select(CONFIG_BASIC_COLUMNS).
		Joins("left join projects on configs.project_id = projects.id").
		Where(mappedQuery).Order(mappedOrderValue).Find(&configs)
	return configs, result.Error
}

func FindConfigList(query map[string]interface{}, pageRequest PageRequest) ([]Config, error) {
	configs := []Config{}
	mappedQuery := mapConfigQuery(query)
	mappedPageRequest := mapConfigPageRequest(pageRequest)
	mappedOrderValue := mappedPageRequest.Order + " " + mappedPageRequest.Direction
	offsetValue := mappedPageRequest.Page * mappedPageRequest.Size
	limitValue := mappedPageRequest.Size
	result := db.MysqlDB.Table("configs").Select(CONFIG_BASIC_COLUMNS).
		Joins("left join projects on configs.project_id = projects.id").
		Where(mappedQuery).Order(mappedOrderValue).Offset(offsetValue).Limit(limitValue).Find(&configs)
	return configs, result.Error
}

func CountConfig(query map[string]interface{}) (int, error) {
	var count int
	result := db.MysqlDB.Model(&Config{}).Where(query).Count(&count)
	return count, result.Error
}

func UpdateConfig(config Config) (Config, error) {
	existingConfig, err := FindConfigById(config.Id)
	if err != nil {
		return config, err
	}
	config.CreateTime = existingConfig.CreateTime
	config.UpdateTime = time.Now()
	result := db.MysqlDB.Save(&config)
	return config, result.Error
}

func DeleteConfigById(id int) error {
	existingConfig, err := FindConfigById(id)
	if err != nil {
		return err
	}
	result := db.MysqlDB.Delete(&existingConfig)
	return result.Error
}
