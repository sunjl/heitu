package api

import (
	"github.com/jishufan/heitu/client/model"
	"github.com/jishufan/heitu/common/constant"
	"github.com/parnurzeal/gorequest"
)

func AddHost(host model.Host) (gorequest.Response, string, []error) {
	return superAgent.Post(address+constant.CreateHostPath).
		Set(constant.AuthToken, token).
		Send(host).
		End()
}

func GetHost(name string) (gorequest.Response, string, []error) {
	s := superAgent.Get(address+constant.GetHostPath).
		Set(constant.AuthToken, token)
	if name != "" {
		s.Param("name", name)
	}
	return s.End()
}

func ListAllHost(groupName, order, direction string) (gorequest.Response, string, []error) {
	s := superAgent.Get(address+constant.ListAllHostPath).
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

func ListHost(groupName, pageStr, sizeStr, order, direction string) (gorequest.Response, string, []error) {
	s := superAgent.Get(address+constant.ListHostPath).
		Set(constant.AuthToken, token)
	if groupName != "" {
		s.Param("groupName", groupName)
	}
	if pageStr != "" {
		s.Param("page", pageStr)
	}
	if sizeStr != "" {
		s.Param("size", sizeStr)
	}
	if order != "" {
		s.Param("order", order)
	}
	if direction != "" {
		s.Param("direction", direction)
	}
	return s.End()
}

func UpdateHost(host model.Host) (gorequest.Response, string, []error) {
	return superAgent.Post(address+constant.UpdateHostPath).
		Set(constant.AuthToken, token).
		Send(host).
		End()
}

func DeleteHost(host model.Host) (gorequest.Response, string, []error) {
	return superAgent.Post(address+constant.DeleteHostPath).
		Set(constant.AuthToken, token).
		Send(host).
		End()
}
