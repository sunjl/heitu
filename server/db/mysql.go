package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jishufan/heitu/common/constant"
	"log"
)

var (
	MysqlDB *gorm.DB
)

func ConnectMysql(dataSourceName string) {
	var err error
	MysqlDB, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		log.Println(constant.ConnectionFailed, err)
	}
}
