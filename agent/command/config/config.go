package config

import (
	"encoding/json"
	"fmt"
	"github.com/jishufan/heitu/agent/terminal"
	"github.com/jishufan/heitu/client/api"
	"github.com/jishufan/heitu/client/model"
	"github.com/spf13/cobra"
	"net/http"
)

var (
	id          int
	projectId   int
	projectName string
	version     string
	environment string
	page        int
	size        int
	order       string
	direction   string
)

var ConfigCmd = &cobra.Command{
	Use: "config",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

var getConfigCmd = &cobra.Command{
	Use: "get",
	Run: func(cmd *cobra.Command, args []string) {
		query := make(map[string]interface{})
		query["id"] = id
		resp, body, errs := api.GetConfig(query)
		if errs != nil {
			fmt.Println(errs)
		}
		switch resp.StatusCode {
		case http.StatusOK:
			config := model.Config{}
			json.Unmarshal([]byte(body), &config)
			terminal.ShowConfig(config)
		default:
			fmt.Println(resp.StatusCode, http.StatusText(resp.StatusCode))
			fmt.Println(body)
		}
	},
}

var listAllConfigCmd = &cobra.Command{
	Use: "listAll",
	Run: func(cmd *cobra.Command, args []string) {
		query := make(map[string]interface{})
		query["id"] = id
		query["projectId"] = projectId
		query["projectName"] = projectName
		query["version"] = version
		query["environment"] = environment
		pageRequest := model.PageRequest{
			Order:     order,
			Direction: direction,
		}
		resp, body, errs := api.ListAllConfig(query, pageRequest)
		if errs != nil {
			fmt.Println(errs)
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var items []model.Config
			if err := json.Unmarshal([]byte(body), &items); err != nil {
				fmt.Println(err)
			}
			terminal.ListAllConfig(items)
		default:
			fmt.Println(resp.StatusCode, http.StatusText(resp.StatusCode))
			fmt.Println(body)
		}
	},
}

var listConfigCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		query := make(map[string]interface{})
		query["id"] = id
		query["projectId"] = projectId
		query["projectName"] = projectName
		query["version"] = version
		query["environment"] = environment
		pageRequest := model.PageRequest{
			Page:      page,
			Size:      size,
			Order:     order,
			Direction: direction,
		}
		resp, body, errs := api.ListConfig(query, pageRequest)
		if errs != nil {
			fmt.Println(errs)
		}
		switch resp.StatusCode {
		case http.StatusOK:
			pageResponse := model.PageResponse{}
			if err := json.Unmarshal([]byte(body), &pageResponse); err != nil {
				fmt.Println(err)
			}
			terminal.ListConfig(pageResponse)
		default:
			fmt.Println(resp.StatusCode, http.StatusText(resp.StatusCode))
			fmt.Println(body)
		}
	},
}

func init() {
	ConfigCmd.AddCommand(getConfigCmd)
	ConfigCmd.AddCommand(listAllConfigCmd)
	ConfigCmd.AddCommand(listConfigCmd)

	getConfigCmd.Flags().IntVarP(&id, "id", "", 0, "id")

	listAllConfigCmd.Flags().IntVarP(&projectId, "projectId", "", 0, "project id")
	listAllConfigCmd.Flags().StringVarP(&projectName, "projectName", "", "", "project name")
	listAllConfigCmd.Flags().StringVarP(&version, "version", "", "", "version")
	listAllConfigCmd.Flags().StringVarP(&environment, "environment", "", "", "environment")
	listAllConfigCmd.Flags().StringVarP(&order, "order", "", "name", "order")
	listAllConfigCmd.Flags().StringVarP(&direction, "direction", "", "asc", "direction")

	listConfigCmd.Flags().IntVarP(&projectId, "projectId", "", 0, "project id")
	listConfigCmd.Flags().StringVarP(&projectName, "projectName", "", "", "project name")
	listConfigCmd.Flags().StringVarP(&version, "version", "", "", "version")
	listConfigCmd.Flags().StringVarP(&environment, "environment", "", "", "environment")
	listConfigCmd.Flags().IntVarP(&page, "page", "", 0, "order")
	listConfigCmd.Flags().IntVarP(&size, "size", "", 10, "direction")
	listConfigCmd.Flags().StringVarP(&order, "order", "", "id", "order")
	listConfigCmd.Flags().StringVarP(&direction, "direction", "", "asc", "direction")
}
