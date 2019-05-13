package model

import (
	"time"
)

type Service struct {
	Id          int       `json:"id"`
	GroupName   string    `json:"groupName"`
	ProjectId   int       `json:"projectId"`
	ProjectName string    `json:"projectName"`
	Name        string    `json:"name"`
	Protocol    string    `json:"protocol"`
	HostName    string    `json:"hostName"`
	IP          string    `json:"ip"`
	Port        int       `json:"port"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
}
