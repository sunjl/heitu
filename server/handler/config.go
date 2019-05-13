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

func CreateConfig() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var config repository.Config
		if err := ctx.Bind(&config); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		createdConfig, err := repository.CreateConfig(config)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusCreated, createdConfig)
	}
}

func GetConfig() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		idStr := ctx.QueryParam("id")
		var (
			id     int
			config repository.Config
			err    error
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
			config, err = repository.FindConfigById(id)
		}
		if err != nil {
			str := err.Error()
			switch {
			case strings.Contains(strings.ToLower(str), constant.RecordNotFound):
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, config)
	}
}

func CountConfig() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		projectIdStr := ctx.QueryParam("projectId")
		projectName := ctx.QueryParam("projectName")
		version := ctx.QueryParam("version")
		environment := ctx.QueryParam("environment")
		var (
			projectId int
			query     = make(map[string]interface{})
			err       error
		)
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
		if version != "" {
			query["version"] = version
		}
		if environment != "" {
			query["environment"] = environment
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		count, err := repository.CountConfig(query)
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

func ListAllConfig() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pageRequest, err := repository.GetPageRequest(ctx)
		projectIdStr := ctx.QueryParam("projectId")
		projectName := ctx.QueryParam("projectName")
		version := ctx.QueryParam("version")
		environment := ctx.QueryParam("environment")
		var (
			projectId int
			query     = make(map[string]interface{})
		)
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
		if version != "" {
			query["version"] = version
		}
		if environment != "" {
			query["environment"] = environment
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		items, err := repository.FindAllConfigList(query, pageRequest)
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

func ListConfig() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pageRequest, err := repository.GetPageRequest(ctx)
		projectIdStr := ctx.QueryParam("projectId")
		projectName := ctx.QueryParam("projectName")
		version := ctx.QueryParam("version")
		environment := ctx.QueryParam("environment")
		var (
			projectId int
			query     = make(map[string]interface{})
		)
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
		if version != "" {
			query["version"] = version
		}
		if environment != "" {
			query["environment"] = environment
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		items, err := repository.FindConfigList(query, pageRequest)
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
		count, err := repository.CountConfig(query)
		pageResponse := repository.PageResponse{
			Page:  pageRequest.Page,
			Size:  pageRequest.Size,
			Count: count,
			Items: util.GetInterfaceSlice(items)}
		return ctx.JSON(http.StatusOK, pageResponse)
	}
}

func UpdateConfig() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var config repository.Config
		if err := ctx.Bind(&config); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if _, err := repository.FindConfigById(config.Id); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		updatedConfig, err := repository.UpdateConfig(config)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, updatedConfig)
	}
}

func DeleteConfig() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var config repository.Config
		if err := ctx.Bind(&config); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if _, err := repository.FindConfigById(config.Id); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		if err := repository.DeleteConfigById(config.Id); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.NoContent(http.StatusNoContent)
	}
}
