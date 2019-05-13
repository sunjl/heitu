package model

import (
	"time"
)

type User struct {
	Id          int       `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	DisplayName string    `json:"displayName"`
	Phone       string    `json:"phone"`
	UserType    string    `json:"userType"`
	Token       string    `json:"token"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
}

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserUpdatePassword struct {
	Username    string `json:"username"`
	NewPassword string `json:"newPassword"`
}
