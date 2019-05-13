package model

import (
	"time"
)

type Host struct {
	Id         int       `json:"id"`
	GroupName  string    `json:"groupName"`
	Name       string    `json:"name"`
	IP         string    `json:"ip"`
	Processor  string    `json:"processor"`
	Memory     int       `json:"memory"`
	Disk       int       `json:"disk"`
	Platform   string    `json:"platform"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}
