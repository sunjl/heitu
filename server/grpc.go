package main

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jishufan/heitu/client/etcd"
	"github.com/jishufan/heitu/common/constant"
	pb "github.com/jishufan/heitu/common/protobuf"
	"github.com/jishufan/heitu/server/db"
	"github.com/jishufan/heitu/server/rpc"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
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
	listener, err := net.Listen("tcp", viper.GetString("grpc.port"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &rpc.UserServiceServer{})
	server.Serve(listener)
}
