package main

import (
	"fmt"
	"github.com/jishufan/heitu/client/api"
	"github.com/jishufan/heitu/client/etcd"
	"github.com/jishufan/heitu/common/constant"
	"github.com/spf13/viper"
)

var (
	cfgFile string
)

func init() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
	viper.AddConfigPath("./config")
	viper.SetConfigName("dev")
	viper.SetConfigType("json")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	api.InitConfig(
		viper.GetString(constant.ApiAddressKey),
		viper.GetString(constant.ApiUserTokenKey),
	)
	etcd.InitConfig(
		viper.GetStringSlice(constant.EtcdEndpointsKey),
		viper.GetInt(constant.EtcdDialTimeoutKey),
		viper.GetInt(constant.EtcdRequestTimeoutKey),
	)
}

func main() {
	etcd.WatchHost()
}
