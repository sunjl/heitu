package repository

import (
	"errors"
	"github.com/jishufan/heitu/common/constant"
	"github.com/jishufan/heitu/common/util"
	"github.com/jishufan/heitu/server/db"
	"time"
)

type User struct {
	Id          int       `gorm:"column:id;primary_key" json:"id"`
	Username    string    `gorm:"column:username" json:"username"`
	Password    string    `gorm:"column:password" json:"password"`
	DisplayName string    `gorm:"column:display_name" json:"displayName"`
	Phone       string    `gorm:"column:phone" json:"phone"`
	UserType    string    `gorm:"column:user_type" json:"userType"`
	Token       string    `gorm:"column:token" json:"token"`
	CreateTime  time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateTime  time.Time `gorm:"column:update_time" json:"updateTime"`
}

func (User) TableName() string {
	return "users"
}

var (
	USER_BASIC_COLUMNS = []string{"id", "username", "display_name", "phone", "user_type", "create_time", "update_time"}
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserUpdatePassword struct {
	Username    string `json:"username"`
	NewPassword string `json:"newPassword"`
}

func CreateUser(user User) (User, error) {
	hashedPassword := util.EncryptPassword(user.Password)
	token := util.UUID()
	user.CreateTime = time.Now()
	user.Password = hashedPassword
	user.Token = token
	result := db.MysqlDB.Create(&user)
	return user, result.Error
}

func FindUserIdByToken(token string) (int, error) {
	user := User{}
	result := db.MysqlDB.Select("id").Where(&User{Token: token}).First(&user)
	return user.Id, result.Error
}

func FindUserTypeByToken(token string) (string, error) {
	user := User{}
	result := db.MysqlDB.Select("user_type").Where(&User{Token: token}).First(&user)
	return user.UserType, result.Error
}

func FindUserBasicById(id int) (User, error) {
	user := User{}
	result := db.MysqlDB.Select(USER_BASIC_COLUMNS).Where(&User{Id: id}).First(&user)
	return user, result.Error
}

func FindUserFullById(id int) (User, error) {
	user := User{}
	result := db.MysqlDB.Where(&User{Id: id}).First(&user)
	if result.RowsAffected == 0 {
		return user, errors.New(constant.RecordNotFound)
	}
	return user, result.Error
}

func FindUserBasicByToken(token string) (User, error) {
	user := User{}
	result := db.MysqlDB.Select(USER_BASIC_COLUMNS).Where(&User{Token: token}).First(&user)
	return user, result.Error
}

func FindUserFullByToken(token string) (User, error) {
	user := User{}
	result := db.MysqlDB.Where(&User{Token: token}).First(&user)
	return user, result.Error
}

func FindUserBasicByUsername(username string) (User, error) {
	user := User{}
	result := db.MysqlDB.Select(USER_BASIC_COLUMNS).Where(&User{Username: username}).First(&user)
	return user, result.Error
}

func FindUserFullByUsername(username string) (User, error) {
	user := User{}
	result := db.MysqlDB.Where(&User{Username: username}).First(&user)
	return user, result.Error
}

func FindAllUserBasicList(query map[string]interface{}, pageRequest PageRequest) ([]User, error) {
	users := []User{}
	orderValue := pageRequest.Order + " " + pageRequest.Direction
	result := db.MysqlDB.Select(USER_BASIC_COLUMNS).Where(query).Order(orderValue).Find(&users)
	return users, result.Error
}

func FindAllUserFullList(query map[string]interface{}, pageRequest PageRequest) ([]User, error) {
	users := []User{}
	orderValue := pageRequest.Order + " " + pageRequest.Direction
	result := db.MysqlDB.Where(query).Order(orderValue).Find(&users)
	return users, result.Error
}

func FindUserBasicList(query map[string]interface{}, pageRequest PageRequest) ([]User, error) {
	users := []User{}
	orderValue := pageRequest.Order + " " + pageRequest.Direction
	offsetValue := pageRequest.Page * pageRequest.Size
	limitValue := pageRequest.Size
	result := db.MysqlDB.Select(USER_BASIC_COLUMNS).Where(query).Order(orderValue).Offset(offsetValue).Limit(limitValue).Find(&users)
	return users, result.Error
}

func FindUserFullList(query map[string]interface{}, pageRequest PageRequest) ([]User, error) {
	users := []User{}
	orderValue := pageRequest.Order + " " + pageRequest.Direction
	offsetValue := pageRequest.Page * pageRequest.Size
	limitValue := pageRequest.Size
	result := db.MysqlDB.Where(query).Order(orderValue).Offset(offsetValue).Limit(limitValue).Find(&users)
	return users, result.Error
}

func CountUser(query map[string]interface{}) (int, error) {
	var count int
	result := db.MysqlDB.Model(&User{}).Where(query).Count(&count)
	return count, result.Error
}

func UpdateUser(user User) (User, error) {
	existingUser, err := FindUserFullById(user.Id)
	if err != nil {
		return user, err
	}
	user.Username = existingUser.Username
	user.Password = existingUser.Password
	user.Token = existingUser.Token
	user.CreateTime = existingUser.CreateTime
	user.UpdateTime = time.Now()
	result := db.MysqlDB.Save(&user)
	return user, result.Error
}

func DeleteUserById(id int) error {
	existingUser, err := FindUserBasicById(id)
	if err != nil {
		return err
	}
	result := db.MysqlDB.Delete(&existingUser)
	return result.Error
}
