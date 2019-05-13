package handler

import (
	"github.com/jishufan/heitu/common/constant"
	"github.com/jishufan/heitu/common/util"
	"github.com/jishufan/heitu/server/repository"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"strings"
)

func CreateService() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var service repository.Service
		if err := ctx.Bind(&service); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		createdService, err := repository.CreateService(service)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusCreated, createdService)
	}
}

func GetService() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		idStr := ctx.QueryParam("id")
		var (
			id      int
			service repository.Service
			err     error
		)
		if idStr == "" {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		switch {
		case idStr != "":
			id, err = strconv.Atoi(idStr)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			service, err = repository.FindServiceById(id)
		}
		if err != nil {
			str := err.Error()
			switch {
			case strings.Contains(strings.ToLower(str), constant.RecordNotFound):
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, service)
	}
}

func CountService() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		groupName := ctx.QueryParam("groupName")
		projectIdStr := ctx.QueryParam("projectId")
		projectName := ctx.QueryParam("projectName")
		protocol := ctx.QueryParam("protocol")
		hostName := ctx.QueryParam("hostName")
		ip := ctx.QueryParam("ip")
		portStr := ctx.QueryParam("port")
		var (
			projectId int
			port      int
			query     = make(map[string]interface{})
			err       error
		)
		if groupName != "" {
			query["group_name"] = groupName
		}
		if projectIdStr != "" {
			projectId, err = strconv.Atoi(projectIdStr)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			query["project_id"] = projectId
		}
		if projectName != "" {
			query["project_name"] = projectName
		}
		if protocol != "" {
			query["protocol"] = protocol
		}
		if hostName != "" {
			query["host_name"] = hostName
		}
		if ip != "" {
			query["ip"] = ip
		}
		if portStr != "" {
			port, err = strconv.Atoi(portStr)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			query["port"] = port
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		count, err := repository.CountService(query)
		if err != nil {
			str := err.Error()
			switch {
			case strings.Contains(strings.ToLower(str), constant.UnknownColumn):
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			case strings.Contains(strings.ToLower(str), constant.RecordNotFound):
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, count)
	}
}

func ListAllService() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pageRequest, err := repository.GetPageRequest(ctx)
		groupName := ctx.QueryParam("groupName")
		projectIdStr := ctx.QueryParam("projectId")
		projectName := ctx.QueryParam("projectName")
		protocol := ctx.QueryParam("protocol")
		hostName := ctx.QueryParam("hostName")
		ip := ctx.QueryParam("ip")
		portStr := ctx.QueryParam("port")
		var (
			projectId int
			port      int
			query     = make(map[string]interface{})
		)
		if groupName != "" {
			query["group_name"] = groupName
		}
		if projectIdStr != "" {
			projectId, err = strconv.Atoi(projectIdStr)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			query["project_id"] = projectId
		}
		if projectName != "" {
			query["project_name"] = projectName
		}
		if protocol != "" {
			query["protocol"] = protocol
		}
		if hostName != "" {
			query["host_name"] = hostName
		}
		if ip != "" {
			query["ip"] = ip
		}
		if portStr != "" {
			port, err = strconv.Atoi(portStr)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			query["port"] = port
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		items, err := repository.FindAllServiceList(query, pageRequest)
		if err != nil {
			str := err.Error()
			switch {
			case strings.Contains(strings.ToLower(str), constant.UnknownColumn):
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			case strings.Contains(strings.ToLower(str), constant.RecordNotFound):
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, items)
	}
}

func ListService() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pageRequest, err := repository.GetPageRequest(ctx)
		groupName := ctx.QueryParam("groupName")
		projectIdStr := ctx.QueryParam("projectId")
		projectName := ctx.QueryParam("projectName")
		protocol := ctx.QueryParam("protocol")
		hostName := ctx.QueryParam("hostName")
		ip := ctx.QueryParam("ip")
		portStr := ctx.QueryParam("port")
		var (
			projectId int
			port      int
			query     = make(map[string]interface{})
		)
		if groupName != "" {
			query["group_name"] = groupName
		}
		if projectIdStr != "" {
			projectId, err = strconv.Atoi(projectIdStr)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			query["project_id"] = projectId
		}
		if projectName != "" {
			query["project_name"] = projectName
		}
		if protocol != "" {
			query["protocol"] = protocol
		}
		if hostName != "" {
			query["host_name"] = hostName
		}
		if ip != "" {
			query["ip"] = ip
		}
		if portStr != "" {
			port, err = strconv.Atoi(portStr)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			query["port"] = port
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		items, err := repository.FindServiceList(query, pageRequest)
		if err != nil {
			str := err.Error()
			switch {
			case strings.Contains(strings.ToLower(str), constant.UnknownColumn):
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			case strings.Contains(strings.ToLower(str), constant.RecordNotFound):
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		count, err := repository.CountService(query)
		pageResponse := repository.PageResponse{
			Page:  pageRequest.Page,
			Size:  pageRequest.Size,
			Count: count,
			Items: util.GetInterfaceSlice(items)}
		return ctx.JSON(http.StatusOK, pageResponse)
	}
}

func UpdateService() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var service repository.Service
		if err := ctx.Bind(&service); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if _, err := repository.FindServiceById(service.Id); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		updatedService, err := repository.UpdateService(service)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, updatedService)
	}
}

func DeleteService() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var service repository.Service
		if err := ctx.Bind(&service); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if _, err := repository.FindServiceById(service.Id); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		if err := repository.DeleteServiceById(service.Id); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.NoContent(http.StatusNoContent)
	}
}
