package terminal

import (
	"fmt"
	"github.com/gosuri/uitable"
	"github.com/jishufan/heitu/client/model"
	"github.com/jishufan/heitu/common/util"
)

func ShowService(service model.Service) {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	table.AddRow("Id:", service.Id)
	table.AddRow("ProjectId:", service.ProjectId)
	table.AddRow("ProjectName:", service.ProjectName)
	table.AddRow("IP:", service.IP)
	table.AddRow("Port:", service.Port)
	table.AddRow("HostName:", service.HostName)
	table.AddRow("CreateTime:", service.CreateTime)
	table.AddRow("UpdateTime:", service.UpdateTime)
	table.AddRow("\n")
	fmt.Printf("--------------------\n")
	fmt.Printf("Service detail: \n")
	fmt.Printf("\n")
	fmt.Println(table)
}

func ListAllService(services []model.Service) {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	for _, service := range services {
		table.AddRow("Id:", service.Id)
		table.AddRow("GroupName:", service.GroupName)
		table.AddRow("ProjectId:", service.ProjectId)
		table.AddRow("ProjectName:", service.ProjectName)
		table.AddRow("Name:", service.Name)
		table.AddRow("Protocol:", service.Protocol)
		table.AddRow("HostName:", service.HostName)
		table.AddRow("IP:", service.IP)
		table.AddRow("Port:", service.Port)
		table.AddRow("CreateTime:", service.CreateTime)
		table.AddRow("UpdateTime:", service.UpdateTime)
		table.AddRow("\n")
	}
	fmt.Printf("--------------------\n")
	fmt.Printf("Total %d services: \n", len(services))
	fmt.Printf("\n")
	fmt.Println(table)
}

func ListService(pageResponse model.PageResponse) {
	count := pageResponse.Count
	items := pageResponse.Items
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	for _, item := range items {
		var service model.Service
		util.DecodeJSON(item, &service)
		table.AddRow("Id:", service.Id)
		table.AddRow("GroupName:", service.GroupName)
		table.AddRow("ProjectId:", service.ProjectId)
		table.AddRow("ProjectName:", service.ProjectName)
		table.AddRow("Name:", service.Name)
		table.AddRow("Protocol:", service.Protocol)
		table.AddRow("HostName:", service.HostName)
		table.AddRow("IP:", service.IP)
		table.AddRow("Port:", service.Port)
		table.AddRow("CreateTime:", service.CreateTime)
		table.AddRow("UpdateTime:", service.UpdateTime)
		table.AddRow("\n")
	}
	fmt.Printf("Total %d services: \n", count)
	fmt.Printf("--------------------\n")
	fmt.Println(table)
}
