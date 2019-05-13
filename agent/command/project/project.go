package project

import (
	"encoding/json"
	"fmt"
	"github.com/jishufan/heitu/agent/terminal"
	"github.com/jishufan/heitu/client/api"
	"github.com/jishufan/heitu/client/model"
	"github.com/spf13/cobra"
	"net/http"
	"strconv"
)

var (
	groupName string
	name      string
	page      int
	size      int
	order     string
	direction string
)

var ProjectCmd = &cobra.Command{
	Use: "project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("project called")
	},
}

var getProjectCmd = &cobra.Command{
	Use: "get",
	Run: func(cmd *cobra.Command, args []string) {
		resp, body, errs := api.GetProject(name)
		if errs != nil {
			fmt.Println(errs)
		}
		switch resp.StatusCode {
		case http.StatusOK:
			project := model.Project{}
			json.Unmarshal([]byte(body), &project)
			terminal.ShowProject(project)
		default:
			fmt.Println(resp.StatusCode, http.StatusText(resp.StatusCode))
			fmt.Println(body)
		}
	},
}

var listAllProjectCmd = &cobra.Command{
	Use: "listAll",
	Run: func(cmd *cobra.Command, args []string) {
		resp, body, errs := api.ListAllProject(groupName, order, direction)
		if errs != nil {
			fmt.Println(errs)
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var items []model.Project
			if err := json.Unmarshal([]byte(body), &items); err != nil {
				fmt.Println(err)
			}
			terminal.ListAllProject(items)
		default:
			fmt.Println(resp.StatusCode, http.StatusText(resp.StatusCode))
			fmt.Println(body)
		}
	},
}

var listProjectCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		pageStr := strconv.Itoa(page)
		sizeStr := strconv.Itoa(size)
		resp, body, errs := api.ListProject(groupName, pageStr, sizeStr, order, direction)
		if errs != nil {
			fmt.Println(errs)
		}
		switch resp.StatusCode {
		case http.StatusOK:
			pageResponse := model.PageResponse{}
			if err := json.Unmarshal([]byte(body), &pageResponse); err != nil {
				fmt.Println(err)
			}
			terminal.ListProject(pageResponse)
		default:
			fmt.Println(resp.StatusCode, http.StatusText(resp.StatusCode))
			fmt.Println(body)
		}
	},
}

func init() {
	ProjectCmd.AddCommand(getProjectCmd)
	ProjectCmd.AddCommand(listAllProjectCmd)
	ProjectCmd.AddCommand(listProjectCmd)

	getProjectCmd.Flags().StringVarP(&name, "name", "", "", "project name")

	listAllProjectCmd.Flags().StringVarP(&groupName, "groupName", "", "", "group name")
	listAllProjectCmd.Flags().StringVarP(&order, "order", "", "name", "order")
	listAllProjectCmd.Flags().StringVarP(&direction, "direction", "", "asc", "direction")

	listProjectCmd.Flags().StringVarP(&groupName, "groupName", "", "", "group name")
	listProjectCmd.Flags().IntVarP(&page, "page", "", 0, "order")
	listProjectCmd.Flags().IntVarP(&size, "size", "", 10, "direction")
	listProjectCmd.Flags().StringVarP(&order, "order", "", "name", "order")
	listProjectCmd.Flags().StringVarP(&direction, "direction", "", "asc", "direction")
}
