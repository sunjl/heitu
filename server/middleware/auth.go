package middleware

import (
	"github.com/jishufan/heitu/common/constant"
	"github.com/jishufan/heitu/server/repository"
	"github.com/jishufan/heitu/server/util"
	"github.com/labstack/echo"
	"net/http"
)

func IsAuthenticated() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(ctx echo.Context) error {
			result, err := util.IsAuthenticated(ctx, repository.FindUserIdByToken)
			if err != nil {
				switch err.Error() {
				case constant.EmptyToken, constant.InvalidToken:
					return echo.NewHTTPError(http.StatusBadRequest, err.Error())
				case constant.RecordNotFound:
					return echo.NewHTTPError(http.StatusUnauthorized)
				}
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			if !result {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}
			return next(ctx)
		})
	}
}

func IsUserType(userType string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(ctx echo.Context) error {
			result, err := util.IsUserType(ctx, userType, repository.FindUserTypeByToken)
			if err != nil {
				switch err.Error() {
				case constant.EmptyToken, constant.InvalidToken:
					return echo.NewHTTPError(http.StatusBadRequest, err.Error())
				case constant.RecordNotFound:
					return echo.NewHTTPError(http.StatusUnauthorized)
				}
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
			if !result {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}
			return next(ctx)
		})
	}
}
