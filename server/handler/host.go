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

func CreateHost() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var host repository.Host
		if err := ctx.Bind(&host); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		createdHost, err := repository.CreateHost(host)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusCreated, createdHost)
	}
}

func GetHost() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		idStr := ctx.QueryParam("id")
		name := ctx.QueryParam("name")
		ip := ctx.QueryParam("ip")
		var (
			id   int
			host repository.Host
			err  error
		)
		if idStr == "" && name == "" && ip == "" {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		switch {
		case idStr != "":
			id, err = strconv.Atoi(idStr)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			host, err = repository.FindHostById(id)
		case name != "":
			host, err = repository.FindHostByName(name)
		case ip != "":
			host, err = repository.FindHostByIP(ip)
		}
		if err != nil {
			str := err.Error()
			switch {
			case strings.Contains(strings.ToLower(str), constant.RecordNotFound):
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, host)
	}
}

func CountHost() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		groupName := ctx.QueryParam("groupName")
		var (
			query = make(map[string]interface{})
		)
		if groupName != "" {
			query["group_name"] = groupName
		}
		count, err := repository.CountHost(query)
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

func ListAllHost() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pageRequest, err := repository.GetPageRequest(ctx)
		groupName := ctx.QueryParam("groupName")
		var (
			query = make(map[string]interface{})
		)
		if groupName != "" {
			query["group_name"] = groupName
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		items, err := repository.FindAllHostList(query, pageRequest)
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

func ListHost() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pageRequest, err := repository.GetPageRequest(ctx)
		groupName := ctx.QueryParam("groupName")
		var (
			query = make(map[string]interface{})
		)
		if groupName != "" {
			query["group_name"] = groupName
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		items, err := repository.FindHostList(query, pageRequest)
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
		count, err := repository.CountHost(query)
		pageResponse := repository.PageResponse{
			Page:  pageRequest.Page,
			Size:  pageRequest.Size,
			Count: count,
			Items: util.GetInterfaceSlice(items)}
		return ctx.JSON(http.StatusOK, pageResponse)
	}
}

func UpdateHost() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var host repository.Host
		if err := ctx.Bind(&host); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if _, err := repository.FindHostById(host.Id); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		updatedHost, err := repository.UpdateHost(host)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, updatedHost)
	}
}

func DeleteHost() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var host repository.Host
		if err := ctx.Bind(&host); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if _, err := repository.FindHostById(host.Id); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		if err := repository.DeleteHostById(host.Id); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.NoContent(http.StatusNoContent)
	}
}
