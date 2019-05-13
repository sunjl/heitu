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

func CreateTask() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var task repository.Task
		if err := ctx.Bind(&task); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		createdTask, err := repository.CreateTask(task)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusCreated, createdTask)
	}
}

func GetTask() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		idStr := ctx.QueryParam("id")
		var (
			id   int
			task repository.Task
			err  error
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
			task, err = repository.FindTaskById(id)
		}
		if err != nil {
			str := err.Error()
			switch {
			case strings.Contains(strings.ToLower(str), constant.RecordNotFound):
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, task)
	}
}

func CountTask() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		hostName := ctx.QueryParam("hostName")
		ip := ctx.QueryParam("ip")
		status := ctx.QueryParam("status")
		var (
			query = make(map[string]interface{})
		)
		if hostName != "" {
			query["host_name"] = hostName
		}
		if ip != "" {
			query["ip"] = ip
		}
		if status != "" {
			query["status"] = status
		}
		count, err := repository.CountTask(query)
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

func ListAllTask() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pageRequest, err := repository.GetPageRequest(ctx)
		hostName := ctx.QueryParam("hostName")
		ip := ctx.QueryParam("ip")
		status := ctx.QueryParam("status")
		var (
			query = make(map[string]interface{})
		)
		if hostName != "" {
			query["host_name"] = hostName
		}
		if ip != "" {
			query["ip"] = ip
		}
		if status != "" {
			query["status"] = status
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		items, err := repository.FindAllTaskList(query, pageRequest)
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

func ListTask() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pageRequest, err := repository.GetPageRequest(ctx)
		hostName := ctx.QueryParam("hostName")
		ip := ctx.QueryParam("ip")
		status := ctx.QueryParam("status")
		var (
			query = make(map[string]interface{})
		)
		if hostName != "" {
			query["host_name"] = hostName
		}
		if ip != "" {
			query["ip"] = ip
		}
		if status != "" {
			query["status"] = status
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		items, err := repository.FindTaskList(query, pageRequest)
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
		count, err := repository.CountTask(query)
		pageResponse := repository.PageResponse{
			Page:  pageRequest.Page,
			Size:  pageRequest.Size,
			Count: count,
			Items: util.GetInterfaceSlice(items)}
		return ctx.JSON(http.StatusOK, pageResponse)
	}
}

func UpdateTask() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var task repository.Task
		if err := ctx.Bind(&task); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if _, err := repository.FindTaskById(task.Id); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		updatedTask, err := repository.UpdateTask(task)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, updatedTask)
	}
}

func DeleteTask() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var task repository.Task
		if err := ctx.Bind(&task); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if _, err := repository.FindTaskById(task.Id); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		if err := repository.DeleteTaskById(task.Id); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.NoContent(http.StatusNoContent)
	}
}
