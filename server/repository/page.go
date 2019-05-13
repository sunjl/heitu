package repository

import (
	"github.com/jishufan/heitu/common/util"
	"github.com/labstack/echo"
	"strconv"
	"strings"
)

type PageRequest struct {
	Page      int    `json:"page"`
	Size      int    `json:"size"`
	Order     string `json:"order"`
	Direction string `json:"direction"`
}

type PageResponse struct {
	Page  int           `json:"page"`
	Size  int           `json:"size"`
	Count int           `json:"count"`
	Items []interface{} `json:"items"`
}

func GetPageRequest(ctx echo.Context) (PageRequest, error) {
	pageParam := ctx.QueryParam("page")
	sizeParam := ctx.QueryParam("size")
	orderParam := ctx.QueryParam("order")
	directionParam := ctx.QueryParam("direction")
	var (
		page      int
		size      int
		order     string
		direction string
		err       error
	)

	switch pageParam {
	case "":
		page = 0
	default:
		page, err = strconv.Atoi(pageParam)
		if err != nil {
			return PageRequest{}, err
		}
	}

	switch sizeParam {
	case "":
		size = 10
	default:
		size, err = strconv.Atoi(sizeParam)
		if err != nil {
			return PageRequest{}, err
		}
	}

	switch orderParam {
	case "":
		order = "id"
	default:
		order = util.ToSnake(orderParam)
	}

	switch directionParam {
	case "":
		direction = "asc"
	default:
		direction = strings.ToLower(directionParam)
	}

	return PageRequest{Page: page, Size: size, Order: order, Direction: direction}, nil
}
