package api

import (
	"github.com/jishufan/heitu/client/model"
	"github.com/jishufan/heitu/common/constant"
	"github.com/parnurzeal/gorequest"
	"strconv"
)

func addConfigQueryParam(s *gorequest.SuperAgent, query map[string]interface{}, pageRequest model.PageRequest) {
	if value, found := query["projectId"]; found {
		if projectId, ok := value.(int); ok {
			s.Param("projectId", strconv.Itoa(projectId))
		}
	}
	if value, found := query["projectName"]; found {
		if projectName, ok := value.(string); ok {
			s.Param("projectName", projectName)
		}
	}
	if value, found := query["version"]; found {
		if version, ok := value.(string); ok {
			s.Param("version", version)
		}
	}
	if value, found := query["environment"]; found {
		if environment, ok := value.(string); ok {
			s.Param("environment", environment)
		}
	}
	s.Param("page", strconv.Itoa(pageRequest.Page))
	s.Param("size", strconv.Itoa(pageRequest.Size))
	s.Param("order", pageRequest.Order)
	s.Param("direction", pageRequest.Direction)
}

func AddConfig(config model.Config) (gorequest.Response, string, []error) {
	return superAgent.Post(address+constant.CreateConfigPath).
		Set(constant.AuthToken, token).
		Send(config).
		End()
}

func GetConfig(query map[string]interface{}) (gorequest.Response, string, []error) {
	s := superAgent.Get(address+constant.GetConfigPath).
		Set(constant.AuthToken, token)
	s.Param("id", strconv.Itoa(query["id"].(int)))
	return s.End()
}

func ListAllConfig(query map[string]interface{}, pageRequest model.PageRequest) (gorequest.Response, string, []error) {
	s := superAgent.Get(address+constant.ListAllConfigPath).
		Set(constant.AuthToken, token)
	addConfigQueryParam(s, query, pageRequest)
	return s.End()
}

func ListConfig(query map[string]interface{}, pageRequest model.PageRequest) (gorequest.Response, string, []error) {
	s := superAgent.Get(address+constant.ListConfigPath).
		Set(constant.AuthToken, token)
	addConfigQueryParam(s, query, pageRequest)
	return s.End()
}

func UpdateConfig(config model.Config) (gorequest.Response, string, []error) {
	return superAgent.Post(address+constant.UpdateConfigPath).
		Set(constant.AuthToken, token).
		Send(config).
		End()
}

func DeleteConfig(config model.Config) (gorequest.Response, string, []error) {
	return superAgent.Post(address+constant.DeleteConfigPath).
		Set(constant.AuthToken, token).
		Send(config).
		End()
}
