package api

import (
	"github.com/jishufan/heitu/client/model"
	"github.com/jishufan/heitu/common/constant"
	"github.com/parnurzeal/gorequest"
	"strconv"
)

func addServiceQueryParam(s *gorequest.SuperAgent, query map[string]interface{}, pageRequest model.PageRequest) {
	if value, found := query["groupName"]; found {
		if groupName, ok := value.(string); ok {
			s.Param("groupName", groupName)
		}
	}
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
	if value, found := query["name"]; found {
		if name, ok := value.(string); ok {
			s.Param("name", name)
		}
	}
	if value, found := query["protocol"]; found {
		if protocol, ok := value.(string); ok {
			s.Param("protocol", protocol)
		}
	}
	if value, found := query["hostName"]; found {
		if hostName, ok := value.(string); ok {
			s.Param("hostName", hostName)
		}
	}
	if value, found := query["ip"]; found {
		if ip, ok := value.(string); ok {
			s.Param("ip", ip)
		}
	}
	if value, found := query["port"]; found {
		if port, ok := value.(int); ok {
			s.Param("port", strconv.Itoa(port))
		}
	}
	s.Param("page", strconv.Itoa(pageRequest.Page))
	s.Param("size", strconv.Itoa(pageRequest.Size))
	s.Param("order", pageRequest.Order)
	s.Param("direction", pageRequest.Direction)
}

func AddService(service model.Service) (gorequest.Response, string, []error) {
	return superAgent.Post(address+constant.CreateServicePath).
		Set(constant.AuthToken, token).
		Send(service).
		End()
}

func GetService(query map[string]interface{}) (gorequest.Response, string, []error) {
	s := superAgent.Get(address+constant.GetServicePath).
		Set(constant.AuthToken, token)
	s.Param("id", strconv.Itoa(query["id"].(int)))
	return s.End()
}

func ListAllService(query map[string]interface{}, pageRequest model.PageRequest) (gorequest.Response, string, []error) {
	s := superAgent.Get(address+constant.ListAllServicePath).
		Set(constant.AuthToken, token)
	addServiceQueryParam(s, query, pageRequest)
	return s.End()
}

func ListService(query map[string]interface{}, pageRequest model.PageRequest) (gorequest.Response, string, []error) {
	s := superAgent.Get(address+constant.ListServicePath).
		Set(constant.AuthToken, token)
	addServiceQueryParam(s, query, pageRequest)
	return s.End()
}

func UpdateService(service model.Service) (gorequest.Response, string, []error) {
	return superAgent.Post(address+constant.UpdateServicePath).
		Set(constant.AuthToken, token).
		Send(service).
		End()
}

func DeleteService(service model.Service) (gorequest.Response, string, []error) {
	return superAgent.Post(address+constant.DeleteServicePath).
		Set(constant.AuthToken, token).
		Send(service).
		End()
}
