package command

import (
	"fmt"
	"github.com/jishufan/heitu/agent/command/config"
	"github.com/jishufan/heitu/agent/command/host"
	"github.com/jishufan/heitu/agent/command/project"
	"github.com/jishufan/heitu/agent/command/service"
	"github.com/jishufan/heitu/client/api"
	"github.com/jishufan/heitu/client/etcd"
	"github.com/jishufan/heitu/common/constant"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile string
)

var RootCmd = &cobra.Command{
	Use: "heitu-agent",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "")
	RootCmd.AddCommand(host.HostCmd)
	RootCmd.AddCommand(project.ProjectCmd)
	RootCmd.AddCommand(config.ConfigCmd)
	RootCmd.AddCommand(service.ServiceCmd)
}

func initConfig() {
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
