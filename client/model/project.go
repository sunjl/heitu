package model

import (
	"time"
)

type Project struct {
	Id         int       `json:"id"`
	GroupName  string    `json:"groupName"`
	Name       string    `json:"name"`
	Git        string    `json:"git"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}
