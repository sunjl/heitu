package model

import (
	"time"
)

type Task struct {
	Id         int       `json:"id"`
	HostName   string    `json:"hostName"`
	IP         string    `json:"ip"`
	Name       string    `json:"name"`
	Content    string    `json:"content"`
	Status     string    `json:"status"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}
