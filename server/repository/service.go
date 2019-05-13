package repository

import (
	"github.com/jishufan/heitu/server/db"
	"time"
)

type Service struct {
	Id          int       `gorm:"column:id;primary_key" json:"id"`
	GroupName   string    `gorm:"column:group_name" json:"groupName"`
	ProjectId   int       `gorm:"column:project_id" json:"projectId"`
	ProjectName string    `gorm:"column:project_name" json:"projectName"`
	Name        string    `gorm:"column:name" json:"name"`
	Protocol    string    `gorm:"column:protocol" json:"protocol"`
	HostName    string    `gorm:"column:host_name" json:"hostName"`
	IP          string    `gorm:"column:ip" json:"ip"`
	Port        int       `gorm:"column:port" json:"port"`
	CreateTime  time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (Service) TableName() string {
	return "services"
}

var (
	SERVICE_COLUMNS = []string{
		"services.id", "services.group_name",
		"services.project_id as project_id", "projects.name as project_name",
		"services.name", "services.protocol",
		"services.host_name", "services.ip", "services.port",
		"services.create_time", "services.update_time",
	}
)

func mapServiceQuery(query map[string]interface{}) map[string]interface{} {
	mappedQuery := make(map[string]interface{})
	for key, value := range query {
		var mappedKey string
		if key == "project_name" {
			mappedKey = "projects.name"
		} else {
			mappedKey = "services." + key
		}
		mappedQuery[mappedKey] = value
	}
	return mappedQuery
}

func mapServicePageRequest(pageRequest PageRequest) PageRequest {
	order := pageRequest.Order
	var mappedOrder string
	if order == "project_name" {
		mappedOrder = "projects.name"
	} else {
		mappedOrder = "services." + order
	}
	return PageRequest{
		Order:     mappedOrder,
		Direction: pageRequest.Direction,
		Page:      pageRequest.Page,
		Size:      pageRequest.Size,
	}
}

func CreateService(service Service) (Service, error) {
	service.CreateTime = time.Now()
	result := db.MysqlDB.Create(&service)
	return service, result.Error
}

func FindServiceById(id int) (Service, error) {
	service := Service{}
	result := db.MysqlDB.Table("services").Select(SERVICE_COLUMNS).
		Joins("left join projects on services.project_id = projects.id").
		Where("services.id = ?", id).First(&service)
	return service, result.Error
}

func FindAllServiceList(query map[string]interface{}, pageRequest PageRequest) ([]Service, error) {
	services := []Service{}
	mappedQuery := mapServiceQuery(query)
	mappedPageRequest := mapServicePageRequest(pageRequest)
	mappedOrderValue := mappedPageRequest.Order + " " + mappedPageRequest.Direction
	result := db.MysqlDB.Table("services").Select(SERVICE_COLUMNS).
		Joins("left join projects on services.project_id = projects.id").
		Where(mappedQuery).Order(mappedOrderValue).Find(&services)
	return services, result.Error
}

func FindServiceList(query map[string]interface{}, pageRequest PageRequest) ([]Service, error) {
	services := []Service{}
	mappedQuery := mapServiceQuery(query)
	mappedPageRequest := mapServicePageRequest(pageRequest)
	mappedOrderValue := mappedPageRequest.Order + " " + mappedPageRequest.Direction
	offsetValue := mappedPageRequest.Page * mappedPageRequest.Size
	limitValue := mappedPageRequest.Size
	result := db.MysqlDB.Table("services").Select(SERVICE_COLUMNS).
		Joins("left join projects on services.project_id = projects.id").
		Where(mappedQuery).Order(mappedOrderValue).Offset(offsetValue).Limit(limitValue).Find(&services)
	return services, result.Error
}

func CountService(query map[string]interface{}) (int, error) {
	var count int
	result := db.MysqlDB.Model(&Service{}).Where(query).Count(&count)
	return count, result.Error
}

func UpdateService(service Service) (Service, error) {
	existingService, err := FindServiceById(service.Id)
	if err != nil {
		return service, err
	}
	service.CreateTime = existingService.CreateTime
	service.UpdateTime = time.Now()
	result := db.MysqlDB.Save(&service)
	return service, result.Error
}

func DeleteServiceById(id int) error {
	existingService, err := FindServiceById(id)
	if err != nil {
		return err
	}
	result := db.MysqlDB.Delete(&existingService)
	return result.Error
}
