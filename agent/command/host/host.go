package host

import (
	"encoding/json"
	"fmt"
	"github.com/jishufan/heitu/agent/shell"
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
	ethernet  string
	disk      string
	page      int
	size      int
	order     string
	direction string
)

var HostCmd = &cobra.Command{
	Use: "host",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("host called")
	},
}

var addHostCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		host := model.Host{
			GroupName: shell.GroupName(groupName),
			Name:      shell.HostName(),
			IP:        shell.IP(ethernet),
			Processor: shell.Processor(),
			Memory:    shell.Memory(),
			Disk:      shell.Disk(disk),
			Platform:  shell.Platform(),
		}
		resp, body, errs := api.AddHost(host)
		if errs != nil {
			fmt.Println(errs)
		}
		switch resp.StatusCode {
		case http.StatusCreated:
			result := model.Host{}
			if err := json.Unmarshal([]byte(body), &result); err != nil {
				fmt.Println(err)
			}
			terminal.ShowHost(result)
		default:
			fmt.Println(resp.StatusCode, http.StatusText(resp.StatusCode))
			fmt.Println(body)
		}
	},
}

var getHostCmd = &cobra.Command{
	Use: "get",
	Run: func(cmd *cobra.Command, args []string) {
		var hostName string
		if name == "" {
			hostName = shell.HostName()
		} else {
			hostName = name
		}
		resp, body, errs := api.GetHost(hostName)
		if errs != nil {
			fmt.Println(errs)
		}
		switch resp.StatusCode {
		case http.StatusOK:
			host := model.Host{}
			json.Unmarshal([]byte(body), &host)
			terminal.ShowHost(host)
		default:
			fmt.Println(resp.StatusCode, http.StatusText(resp.StatusCode))
			fmt.Println(body)
		}
	},
}

var listAllHostCmd = &cobra.Command{
	Use: "listAll",
	Run: func(cmd *cobra.Command, args []string) {
		resp, body, errs := api.ListAllHost(groupName, order, direction)
		if errs != nil {
			fmt.Println(errs)
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var items []model.Host
			if err := json.Unmarshal([]byte(body), &items); err != nil {
				fmt.Println(err)
			}
			terminal.ListAllHost(items)
		default:
			fmt.Println(resp.StatusCode, http.StatusText(resp.StatusCode))
			fmt.Println(body)
		}
	},
}

var listHostCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		pageStr := strconv.Itoa(page)
		sizeStr := strconv.Itoa(size)
		resp, body, errs := api.ListHost(groupName, pageStr, sizeStr, order, direction)
		if errs != nil {
			fmt.Println(errs)
		}
		switch resp.StatusCode {
		case http.StatusOK:
			pageResponse := model.PageResponse{}
			if err := json.Unmarshal([]byte(body), &pageResponse); err != nil {
				fmt.Println(err)
			}
			terminal.ListHost(pageResponse)
		default:
			fmt.Println(resp.StatusCode, http.StatusText(resp.StatusCode))
			fmt.Println(body)
		}
	},
}

var updateHostCmd = &cobra.Command{
	Use: "update",
	Run: func(cmd *cobra.Command, args []string) {
		resp, body, errs := api.GetHost(shell.HostName())
		if errs != nil {
			fmt.Println(errs)
		}
		switch resp.StatusCode {
		case http.StatusOK:
			existingHost := model.Host{}
			json.Unmarshal([]byte(body), &existingHost)
			hostToUpdate := model.Host{
				Id:        existingHost.Id,
				GroupName: shell.GroupName(groupName),
				Name:      shell.HostName(),
				IP:        shell.IP(ethernet),
				Processor: shell.Processor(),
				Memory:    shell.Memory(),
				Disk:      shell.Disk(disk),
				Platform:  shell.Platform(),
			}
			resp, body, errs := api.UpdateHost(hostToUpdate)
			if errs != nil {
				fmt.Println(errs)
			}
			switch resp.StatusCode {
			case http.StatusOK:
				result := model.Host{}
				if err := json.Unmarshal([]byte(body), &result); err != nil {
					fmt.Println(err)
				}
				terminal.ShowHost(result)
			default:
				fmt.Println(resp.StatusCode, http.StatusText(resp.StatusCode))
				fmt.Println(body)
			}
		default:
			fmt.Println(resp.StatusCode, http.StatusText(resp.StatusCode))
			fmt.Println(body)
		}
	},
}

var deleteHostCmd = &cobra.Command{
	Use: "delete",
	Run: func(cmd *cobra.Command, args []string) {
		var hostName string
		if name == "" {
			hostName = shell.HostName()
		} else {
			hostName = name
		}
		resp, body, errs := api.GetHost(hostName)
		if errs != nil {
			fmt.Println(errs)
		}
		switch resp.StatusCode {
		case http.StatusOK:
			existingHost := model.Host{}
			json.Unmarshal([]byte(body), &existingHost)
			hostToDelete := model.Host{
				Id: existingHost.Id,
			}
			resp, body, errs := api.DeleteHost(hostToDelete)
			if errs != nil {
				fmt.Println(errs)
			}
			switch resp.StatusCode {
			case http.StatusNoContent:
				fmt.Println("Successfully delete host. \n")
			default:
				fmt.Println(resp.StatusCode, http.StatusText(resp.StatusCode))
				fmt.Println(body)
			}
		default:
			fmt.Println(resp.StatusCode, http.StatusText(resp.StatusCode))
			fmt.Println(body)
		}
	},
}

func init() {
	HostCmd.AddCommand(addHostCmd)
	HostCmd.AddCommand(getHostCmd)
	HostCmd.AddCommand(listAllHostCmd)
	HostCmd.AddCommand(listHostCmd)
	HostCmd.AddCommand(updateHostCmd)
	HostCmd.AddCommand(deleteHostCmd)

	addHostCmd.Flags().StringVarP(&groupName, "groupName", "", "", "group name")
	addHostCmd.Flags().StringVarP(&ethernet, "ethernet", "", "eth0", "ethernet device")
	addHostCmd.Flags().StringVarP(&disk, "disk", "", "", "disk device")

	getHostCmd.Flags().StringVarP(&name, "name", "", "", "host name")

	listAllHostCmd.Flags().StringVarP(&groupName, "groupName", "", "", "group name")
	listAllHostCmd.Flags().StringVarP(&order, "order", "", "name", "order")
	listAllHostCmd.Flags().StringVarP(&direction, "direction", "", "asc", "direction")

	listHostCmd.Flags().StringVarP(&groupName, "groupName", "", "", "group name")
	listHostCmd.Flags().IntVarP(&page, "page", "", 0, "order")
	listHostCmd.Flags().IntVarP(&size, "size", "", 10, "direction")
	listHostCmd.Flags().StringVarP(&order, "order", "", "name", "order")
	listHostCmd.Flags().StringVarP(&direction, "direction", "", "asc", "direction")

	updateHostCmd.Flags().StringVarP(&groupName, "groupName", "", "", "group name")
	updateHostCmd.Flags().StringVarP(&ethernet, "ethernet", "", "eth0", "ethernet device")
	updateHostCmd.Flags().StringVarP(&disk, "disk", "", "", "disk device")

	deleteHostCmd.Flags().StringVarP(&name, "name", "", "", "host name")
}
