package util

import (
	"errors"
	"github.com/jishufan/heitu/common/constant"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func ComparePassword(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GetToken(ctx echo.Context) (string, error) {
	token := ctx.Request().Header().Get(constant.AuthToken)
	if token == "" {
		return "", errors.New(constant.EmptyToken)
	}
	_, err := uuid.FromString(token)
	if err != nil {
		return "", errors.New(constant.InvalidToken)
	}
	return token, nil
}

func IsAuthenticated(ctx echo.Context, findUserIdByToken func(string) (int, error)) (bool, error) {
	token, err := GetToken(ctx)
	if err != nil {
		return false, err
	}
	if _, err := findUserIdByToken(token); err != nil {
		return false, err
	}
	return true, nil
}

func IsUserType(ctx echo.Context, userType string, findUserTypeByToken func(string) (string, error)) (bool, error) {
	token, err := GetToken(ctx)
	if err != nil {
		return false, err
	}
	if userType == "" {
		return false, nil
	}
	existingUserType, err := findUserTypeByToken(token)
	if err != nil {
		return false, err
	}
	if userType != existingUserType {
		return false, nil
	}
	return true, nil
}
