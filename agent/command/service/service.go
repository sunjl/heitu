package service

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
	groupName   string
	projectId   int
	projectName string
	name        string
	protocol    string
	hostName    string
	ip          string
	port        int
	page        int
	size        int
	order       string
	direction   string
)

var ServiceCmd = &cobra.Command{
	Use: "service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("service called")
	},
}

var getServiceCmd = &cobra.Command{
	Use: "get",
	Run: func(cmd *cobra.Command, args []string) {
		query := make(map[string]interface{})
		query["id"] = id
		resp, body, errs := api.GetService(query)
		if errs != nil {
			fmt.Println(errs)
		}
		switch resp.StatusCode {
		case http.StatusOK:
			service := model.Service{}
			json.Unmarshal([]byte(body), &service)
			terminal.ShowService(service)
		default:
			fmt.Println(resp.StatusCode, http.StatusText(resp.StatusCode))
			fmt.Println(body)
		}
	},
}

var listAllServiceCmd = &cobra.Command{
	Use: "listAll",
	Run: func(cmd *cobra.Command, args []string) {
		query := make(map[string]interface{})
		query["id"] = id
		query["groupName"] = groupName
		query["projectId"] = projectId
		query["projectName"] = projectName
		query["name"] = name
		query["protocol"] = protocol
		query["hostName"] = hostName
		query["ip"] = ip
		query["port"] = port
		pageRequest := model.PageRequest{
			Order:     order,
			Direction: direction,
		}
		resp, body, errs := api.ListAllService(query, pageRequest)
		if errs != nil {
			fmt.Println(errs)
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var items []model.Service
			if err := json.Unmarshal([]byte(body), &items); err != nil {
				fmt.Println(err)
			}
			terminal.ListAllService(items)
		default:
			fmt.Println(resp.StatusCode, http.StatusText(resp.StatusCode))
			fmt.Println(body)
		}
	},
}

var listServiceCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		query := make(map[string]interface{})
		query["id"] = id
		query["groupName"] = groupName
		query["projectId"] = projectId
		query["projectName"] = projectName
		query["name"] = name
		query["protocol"] = protocol
		query["hostName"] = hostName
		query["ip"] = ip
		query["port"] = port
		pageRequest := model.PageRequest{
			Page:      page,
			Size:      size,
			Order:     order,
			Direction: direction,
		}
		resp, body, errs := api.ListService(query, pageRequest)
		if errs != nil {
			fmt.Println(errs)
		}
		switch resp.StatusCode {
		case http.StatusOK:
			pageResponse := model.PageResponse{}
			if err := json.Unmarshal([]byte(body), &pageResponse); err != nil {
				fmt.Println(err)
			}
			terminal.ListService(pageResponse)
		default:
			fmt.Println(resp.StatusCode, http.StatusText(resp.StatusCode))
			fmt.Println(body)
		}
	},
}

func init() {
	ServiceCmd.AddCommand(getServiceCmd)
	ServiceCmd.AddCommand(listAllServiceCmd)
	ServiceCmd.AddCommand(listServiceCmd)

	getServiceCmd.Flags().IntVarP(&id, "id", "", 0, "id")

	listAllServiceCmd.Flags().StringVarP(&groupName, "groupName", "", "", "group name")
	listAllServiceCmd.Flags().IntVarP(&projectId, "projectId", "", 0, "project id")
	listAllServiceCmd.Flags().StringVarP(&projectName, "projectName", "", "", "project name")
	listAllServiceCmd.Flags().StringVarP(&name, "name", "", "", "name")
	listAllServiceCmd.Flags().StringVarP(&protocol, "protocol", "", "", "protocol")
	listAllServiceCmd.Flags().StringVarP(&hostName, "hostName", "", "", "host name")
	listAllServiceCmd.Flags().StringVarP(&ip, "ip", "", "", "ip")
	listAllServiceCmd.Flags().IntVarP(&port, "port", "", 0, "port")
	listAllServiceCmd.Flags().StringVarP(&order, "order", "", "name", "order")
	listAllServiceCmd.Flags().StringVarP(&direction, "direction", "", "asc", "direction")

	listServiceCmd.Flags().StringVarP(&groupName, "groupName", "", "", "group name")
	listServiceCmd.Flags().IntVarP(&projectId, "projectId", "", 0, "project id")
	listServiceCmd.Flags().StringVarP(&projectName, "projectName", "", "", "project name")
	listServiceCmd.Flags().StringVarP(&name, "name", "", "", "name")
	listServiceCmd.Flags().StringVarP(&protocol, "protocol", "", "", "protocol")
	listServiceCmd.Flags().StringVarP(&hostName, "hostName", "", "", "host name")
	listServiceCmd.Flags().StringVarP(&ip, "ip", "", "", "ip")
	listServiceCmd.Flags().IntVarP(&port, "port", "", 0, "port")
	listServiceCmd.Flags().IntVarP(&page, "page", "", 0, "order")
	listServiceCmd.Flags().IntVarP(&size, "size", "", 10, "direction")
	listServiceCmd.Flags().StringVarP(&order, "order", "", "id", "order")
	listServiceCmd.Flags().StringVarP(&direction, "direction", "", "asc", "direction")
}
