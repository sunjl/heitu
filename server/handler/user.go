package handler

import (
	"github.com/jishufan/heitu/common/constant"
	commonUtil "github.com/jishufan/heitu/common/util"
	"github.com/jishufan/heitu/server/repository"
	"github.com/jishufan/heitu/server/util"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"strings"
)

func CreateUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var user repository.User
		if err := ctx.Bind(&user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		createdUser, err := repository.CreateUser(user)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusCreated, createdUser)
	}
}

func SignInUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var userCredentials repository.UserCredentials
		if err := ctx.Bind(&userCredentials); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		existingUser, err := repository.FindUserFullByUsername(userCredentials.Username)
		if err != nil {
			str := err.Error()
			switch {
			case strings.Contains(strings.ToLower(str), constant.RecordNotFound):
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if err := util.ComparePassword(existingUser.Password, userCredentials.Password); err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		return ctx.JSON(http.StatusOK, existingUser)
	}
}

func GetUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		idStr := ctx.QueryParam("id")
		username := ctx.QueryParam("username")
		token := ctx.QueryParam("token")
		var (
			authToken    string
			authUserType string
			id           int
			user         repository.User
			err          error
		)
		if idStr == "" && username == "" && token == "" {
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		authToken, err = util.GetToken(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		authUserType, err = repository.FindUserTypeByToken(authToken)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		switch authUserType {
		case constant.Admin:
			switch {
			case idStr != "":
				id, err = strconv.Atoi(idStr)
				if err != nil {
					return echo.NewHTTPError(http.StatusBadRequest, err.Error())
				}
				user, err = repository.FindUserFullById(id)
			case username != "":
				user, err = repository.FindUserFullByUsername(username)
			case token != "":
				user, err = repository.FindUserFullByToken(token)
			}
		case constant.Normal:
			switch {
			case idStr != "":
				id, err := strconv.Atoi(idStr)
				if err != nil {
					return echo.NewHTTPError(http.StatusBadRequest, err.Error())
				}
				user, err = repository.FindUserBasicById(id)
			case username != "":
				user, err = repository.FindUserBasicByUsername(username)
			case token != "":
				user, err = repository.FindUserBasicByToken(token)
			}
		}
		if err != nil {
			str := err.Error()
			switch {
			case strings.Contains(strings.ToLower(str), constant.RecordNotFound):
				return echo.NewHTTPError(http.StatusNotFound, err.Error())
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, user)
	}
}

func CountUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userType := ctx.QueryParam("userType")
		var (
			query = make(map[string]interface{})
		)
		if userType != "" {
			query["user_type"] = userType
		}
		token, err := util.GetToken(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if _, err := repository.FindUserTypeByToken(token); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		count, err := repository.CountUser(query)
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

func ListAllUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pageRequest, err := repository.GetPageRequest(ctx)
		userType := ctx.QueryParam("userType")
		var (
			query = make(map[string]interface{})
		)
		if userType != "" {
			query["user_type"] = userType
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		token, err := util.GetToken(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		authUserType, err := repository.FindUserTypeByToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		var items []repository.User
		switch authUserType {
		case constant.Admin:
			items, err = repository.FindAllUserFullList(query, pageRequest)
		case constant.Normal:
			items, err = repository.FindAllUserBasicList(query, pageRequest)
		}
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

func ListUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		pageRequest, err := repository.GetPageRequest(ctx)
		userType := ctx.QueryParam("userType")
		var (
			query = make(map[string]interface{})
		)
		if userType != "" {
			query["user_type"] = userType
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		token, err := util.GetToken(ctx)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		authUserType, err := repository.FindUserTypeByToken(token)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		var items []repository.User
		switch authUserType {
		case constant.Admin:
			items, err = repository.FindUserFullList(query, pageRequest)
		case constant.Normal:
			items, err = repository.FindUserBasicList(query, pageRequest)
		}
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
		count, err := repository.CountUser(query)
		pageResponse := repository.PageResponse{
			Page:  pageRequest.Page,
			Size:  pageRequest.Size,
			Count: count,
			Items: commonUtil.GetInterfaceSlice(items)}
		return ctx.JSON(http.StatusOK, pageResponse)
	}
}

func UpdateUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var user repository.User
		if err := ctx.Bind(&user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if _, err := repository.FindUserBasicById(user.Id); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		updatedUser, err := repository.UpdateUser(user)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, updatedUser)
	}
}

func DeleteUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var user repository.User
		if err := ctx.Bind(&user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if _, err := repository.FindUserBasicById(user.Id); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		if err := repository.DeleteUserById(user.Id); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return ctx.NoContent(http.StatusNoContent)
	}
}
