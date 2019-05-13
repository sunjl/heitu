package repository

import (
	"github.com/jishufan/heitu/client/etcd"
	"github.com/jishufan/heitu/client/model"
	"github.com/jishufan/heitu/server/db"
	"time"
)

type Host struct {
	Id         int       `gorm:"column:id;primary_key" json:"id"`
	GroupName  string    `gorm:"column:group_name" json:"groupName"`
	Name       string    `gorm:"column:name" json:"name"`
	IP         string    `gorm:"column:ip" json:"ip"`
	Processor  string    `gorm:"column:processor" json:"processor"`
	Memory     int       `gorm:"column:memory" json:"memory"`
	Disk       int       `gorm:"column:disk" json:"disk"`
	Platform   string    `gorm:"column:platform" json:"platform"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (Host) TableName() string {
	return "hosts"
}

func CreateHost(host Host) (Host, error) {
	host.CreateTime = time.Now()
	result := db.MysqlDB.Create(&host)
	if result.Error != nil {
		return host, result.Error
	}
	etcd.AddHost(model.Host{
		Id:         host.Id,
		GroupName:  host.GroupName,
		Name:       host.Name,
		IP:         host.IP,
		Processor:  host.Processor,
		Memory:     host.Memory,
		Disk:       host.Disk,
		Platform:   host.Platform,
		CreateTime: host.CreateTime,
		UpdateTime: host.UpdateTime,
	})
	return host, nil
}

func FindHostById(id int) (Host, error) {
	host := Host{}
	result := db.MysqlDB.Where(&Host{Id: id}).First(&host)
	return host, result.Error
}

func FindHostByName(name string) (Host, error) {
	host := Host{}
	result := db.MysqlDB.Where(&Host{Name: name}).First(&host)
	return host, result.Error
}

func FindHostByIP(ip string) (Host, error) {
	host := Host{}
	result := db.MysqlDB.Where(&Host{IP: ip}).First(&host)
	return host, result.Error
}

func FindAllHostList(query map[string]interface{}, pageRequest PageRequest) ([]Host, error) {
	hosts := []Host{}
	orderValue := pageRequest.Order + " " + pageRequest.Direction
	result := db.MysqlDB.Where(query).Order(orderValue).Find(&hosts)
	return hosts, result.Error
}

func FindHostList(query map[string]interface{}, pageRequest PageRequest) ([]Host, error) {
	hosts := []Host{}
	orderValue := pageRequest.Order + " " + pageRequest.Direction
	offsetValue := pageRequest.Page * pageRequest.Size
	limitValue := pageRequest.Size
	result := db.MysqlDB.Where(query).Order(orderValue).Offset(offsetValue).Limit(limitValue).Find(&hosts)
	return hosts, result.Error
}

func CountHost(query map[string]interface{}) (int, error) {
	var count int
	result := db.MysqlDB.Model(&Host{}).Where(query).Count(&count)
	return count, result.Error
}

func UpdateHost(host Host) (Host, error) {
	existingHost, err := FindHostById(host.Id)
	if err != nil {
		return host, err
	}
	host.CreateTime = existingHost.CreateTime
	host.UpdateTime = time.Now()
	result := db.MysqlDB.Save(&host)
	return host, result.Error
}

func DeleteHostById(id int) error {
	existingHost, err := FindHostById(id)
	if err != nil {
		return err
	}
	result := db.MysqlDB.Delete(&existingHost)
	return result.Error
}
