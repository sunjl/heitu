package api

import (
	"github.com/jishufan/heitu/client/model"
	"github.com/jishufan/heitu/common/constant"
	"github.com/parnurzeal/gorequest"
)

func AddProject(project model.Project) (gorequest.Response, string, []error) {
	return superAgent.Post(address+constant.CreateProjectPath).
		Set(constant.AuthToken, token).
		Send(project).
		End()
}

func GetProject(name string) (gorequest.Response, string, []error) {
	s := superAgent.Get(address+constant.GetProjectPath).
		Set(constant.AuthToken, token)
	if name != "" {
		s.Param("name", name)
	}
	return s.End()
}

func ListAllProject(groupName, order, direction string) (gorequest.Response, string, []error) {
	s := superAgent.Get(address+constant.ListAllProjectPath).
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

func ListProject(groupName, pageStr, sizeStr, order, direction string) (gorequest.Response, string, []error) {
	s := superAgent.Get(address+constant.ListProjectPath).
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

func UpdateProject(project model.Project) (gorequest.Response, string, []error) {
	return superAgent.Post(address+constant.UpdateProjectPath).
		Set(constant.AuthToken, token).
		Send(project).
		End()
}

func DeleteProject(project model.Project) (gorequest.Response, string, []error) {
	return superAgent.Post(address+constant.DeleteProjectPath).
		Set(constant.AuthToken, token).
		Send(project).
		End()
}
