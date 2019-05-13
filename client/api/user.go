package api

import (
	"github.com/jishufan/heitu/client/model"
	"github.com/jishufan/heitu/common/constant"
	"github.com/parnurzeal/gorequest"
)

func AddUser(user model.User) (gorequest.Response, string, []error) {
	return superAgent.Post(address+constant.CreateUserPath).
		Set(constant.AuthToken, token).
		Send(user).
		End()
}

func GetUser(name string) (gorequest.Response, string, []error) {
	s := superAgent.Get(address+constant.GetUserPath).
		Set(constant.AuthToken, token)
	if name != "" {
		s.Param("name", name)
	}
	return s.End()
}

func ListAllUser(groupName, order, direction string) (gorequest.Response, string, []error) {
	s := superAgent.Get(address+constant.ListAllUserPath).
		Set(constant.AuthToken, token)
	if groupName != "" {
		s.Param("groupName", groupName)
	}
	if order != "" {
		s.Param("order", order)
	}
	if direction != "" {
		s.Param("direction", direction)
	}
	return s.End()
}

func ListUser(groupName, page, size, order, direction string) (gorequest.Response, string, []error) {
	s := superAgent.Get(address+constant.ListUserPath).
		Set(constant.AuthToken, token)
	if groupName != "" {
		s.Param("groupName", groupName)
	}
	if page != "" {
		s.Param("page", page)
	}
	if size != "" {
		s.Param("size", size)
	}
	if order != "" {
		s.Param("order", order)
	}
	if direction != "" {
		s.Param("direction", direction)
	}
	return s.End()
}

func UpdateUser(user model.User) (gorequest.Response, string, []error) {
	return superAgent.Post(address+constant.UpdateUserPath).
		Set(constant.AuthToken, token).
		Send(user).
		End()
}

func DeleteUser(user model.User) (gorequest.Response, string, []error) {
	return superAgent.Post(address+constant.DeleteUserPath).
		Set(constant.AuthToken, token).
		Send(user).
		End()
}
