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

func CreateProject() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var project repository.Project
		if err := ctx.Bind(&project); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		createdProject, err := repository.CreateProject(project)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusCreated, createdProject)
	}
}

func GetProject() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		idStr := ctx.QueryParam("id")
		name := ctx.QueryParam("name")
		var (
			id      int
			project repository.Project
			err     error
		)
		if idStr == "" && name == "" {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		switch {
		case idStr != "":
			id, err = strconv.Atoi(idStr)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			project, err = repository.FindProjectById(id)
		case name != "":
			project, err = repository.FindProjectByName(name)
		}
		if err != nil {
			str := err.Error()
			switch {
			case strings.Contains(strings.ToLower(str), constant.RecordNotFound):
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, project)
	}
}

func CountProject() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		groupName := ctx.QueryParam("groupName")
		var (
			query = make(map[string]interface{})
		)
		if groupName != "" {
			query["group_name"] = groupName
		}
		count, err := repository.CountProject(query)
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

func ListAllProject() echo.HandlerFunc {
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
		items, err := repository.FindAllProjectList(query, pageRequest)
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

func ListProject() echo.HandlerFunc {
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
		items, err := repository.FindProjectList(query, pageRequest)
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
		count, err := repository.CountProject(query)
		pageResponse := repository.PageResponse{
			Page:  pageRequest.Page,
			Size:  pageRequest.Size,
			Count: count,
			Items: util.GetInterfaceSlice(items)}
		return ctx.JSON(http.StatusOK, pageResponse)
	}
}

func UpdateProject() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var project repository.Project
		if err := ctx.Bind(&project); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if _, err := repository.FindProjectById(project.Id); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		updatedProject, err := repository.UpdateProject(project)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, updatedProject)
	}
}

func DeleteProject() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var project repository.Project
		if err := ctx.Bind(&project); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if _, err := repository.FindProjectById(project.Id); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		if err := repository.DeleteProjectById(project.Id); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.NoContent(http.StatusNoContent)
	}
}
