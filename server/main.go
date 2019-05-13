package main

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jishufan/heitu/client/etcd"
	"github.com/jishufan/heitu/common/constant"
	"github.com/jishufan/heitu/server/db"
	"github.com/jishufan/heitu/server/handler"
	middleware "github.com/jishufan/heitu/server/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	echoMiddleware "github.com/labstack/echo/middleware"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"github.com/tylerb/graceful"
	"time"
)

func init() {
	viper.AddConfigPath("./config")
	viper.SetConfigName("dev")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	mysqlConfig := mysql.Config{
		User:      viper.GetString("mysql.username"),
		Passwd:    viper.GetString("mysql.password"),
		Net:       "tcp",
		Addr:      viper.GetString("mysql.address"),
		DBName:    viper.GetString("mysql.db"),
		Loc:       time.Local,
		ParseTime: true,
		Params: map[string]string{
			"charset": "utf8",
		},
	}
	db.ConnectMysql(mysqlConfig.FormatDSN())
	etcd.InitConfig(
		viper.GetStringSlice(constant.EtcdEndpointsKey),
		viper.GetInt(constant.EtcdDialTimeoutKey),
		viper.GetInt(constant.EtcdRequestTimeoutKey),
	)
}

func main() {

	echo := echo.New()

	echo.Use(echoMiddleware.Logger())
	echo.Use(echoMiddleware.Recover())
	echo.Use(echoMiddleware.Gzip())

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Origin", "Accept", "Content-Type", "X-Auth-Token"},
	})
	echo.Use(standard.WrapMiddleware(cors.Handler))

	echo.Post(constant.CreateUserPath, handler.CreateUser(), middleware.IsUserType(constant.Admin))
	echo.Post(constant.SignInUserPath, handler.SignInUser())
	echo.Get(constant.GetUserPath, handler.GetUser(), middleware.IsAuthenticated())
	echo.Get(constant.CountUserPath, handler.CountUser(), middleware.IsAuthenticated())
	echo.Get(constant.ListAllUserPath, handler.ListAllUser(), middleware.IsAuthenticated())
	echo.Get(constant.ListUserPath, handler.ListUser(), middleware.IsAuthenticated())
	echo.Post(constant.UpdateUserPath, handler.UpdateUser(), middleware.IsUserType(constant.Admin))
	echo.Post(constant.DeleteUserPath, handler.DeleteUser(), middleware.IsUserType(constant.Admin))

	echo.Post(constant.CreateHostPath, handler.CreateHost(), middleware.IsAuthenticated())
	echo.Get(constant.GetHostPath, handler.GetHost(), middleware.IsAuthenticated())
	echo.Get(constant.CountHostPath, handler.CountHost(), middleware.IsAuthenticated())
	echo.Get(constant.ListAllHostPath, handler.ListAllHost(), middleware.IsAuthenticated())
	echo.Get(constant.ListHostPath, handler.ListHost(), middleware.IsAuthenticated())
	echo.Post(constant.UpdateHostPath, handler.UpdateHost(), middleware.IsAuthenticated())
	echo.Post(constant.DeleteHostPath, handler.DeleteHost(), middleware.IsUserType(constant.Admin))

	echo.Post(constant.CreateProjectPath, handler.CreateProject(), middleware.IsUserType(constant.Admin))
	echo.Get(constant.GetProjectPath, handler.GetProject(), middleware.IsAuthenticated())
	echo.Get(constant.CountProjectPath, handler.CountProject(), middleware.IsAuthenticated())
	echo.Get(constant.ListAllProjectPath, handler.ListAllProject(), middleware.IsAuthenticated())
	echo.Get(constant.ListProjectPath, handler.ListProject(), middleware.IsAuthenticated())
	echo.Post(constant.UpdateProjectPath, handler.UpdateProject(), middleware.IsUserType(constant.Admin))
	echo.Post(constant.DeleteProjectPath, handler.DeleteProject(), middleware.IsUserType(constant.Admin))

	echo.Post(constant.CreateConfigPath, handler.CreateConfig(), middleware.IsUserType(constant.Admin))
	echo.Get(constant.GetConfigPath, handler.GetConfig(), middleware.IsAuthenticated())
	echo.Get(constant.CountConfigPath, handler.CountConfig(), middleware.IsAuthenticated())
	echo.Get(constant.ListAllConfigPath, handler.ListAllConfig(), middleware.IsAuthenticated())
	echo.Get(constant.ListConfigPath, handler.ListConfig(), middleware.IsAuthenticated())
	echo.Post(constant.UpdateConfigPath, handler.UpdateConfig(), middleware.IsUserType(constant.Admin))
	echo.Post(constant.DeleteConfigPath, handler.DeleteConfig(), middleware.IsUserType(constant.Admin))

	echo.Post(constant.CreateTaskPath, handler.CreateTask(), middleware.IsUserType(constant.Admin))
	echo.Get(constant.GetTaskPath, handler.GetTask(), middleware.IsAuthenticated())
	echo.Get(constant.CountTaskPath, handler.CountTask(), middleware.IsAuthenticated())
	echo.Get(constant.ListAllTaskPath, handler.ListAllTask(), middleware.IsAuthenticated())
	echo.Get(constant.ListTaskPath, handler.ListTask(), middleware.IsAuthenticated())
	echo.Post(constant.UpdateTaskPath, handler.UpdateTask(), middleware.IsAuthenticated())
	echo.Post(constant.DeleteTaskPath, handler.DeleteTask(), middleware.IsUserType(constant.Admin))

	echo.Post(constant.CreateServicePath, handler.CreateService(), middleware.IsUserType(constant.Admin))
	echo.Get(constant.GetServicePath, handler.GetService(), middleware.IsAuthenticated())
	echo.Get(constant.CountServicePath, handler.CountService(), middleware.IsAuthenticated())
	echo.Get(constant.ListAllServicePath, handler.ListAllService(), middleware.IsAuthenticated())
	echo.Get(constant.ListServicePath, handler.ListService(), middleware.IsAuthenticated())
	echo.Post(constant.UpdateServicePath, handler.UpdateService(), middleware.IsUserType(constant.Admin))
	echo.Post(constant.DeleteServicePath, handler.DeleteService(), middleware.IsUserType(constant.Admin))

	std := standard.New(":9000")
	std.SetHandler(echo)
	graceful.ListenAndServe(std.Server, 5*time.Second)

}
