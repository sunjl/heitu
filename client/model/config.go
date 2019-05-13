package model

import (
	"time"
)

type Config struct {
	Id          int       `json:"id"`
	ProjectId   int       `json:"projectId"`
	ProjectName string    `json:"projectName"`
	Version     string    `json:"version"`
	Environment string    `json:"environment"`
	FileName    string    `json:"fileName"`
	Content     string    `json:"content"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
}
