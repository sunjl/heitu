package terminal

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/gosuri/uitable"
	"github.com/jishufan/heitu/client/model"
	"github.com/jishufan/heitu/common/util"
)

func ShowHost(host model.Host) {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	table.AddRow("Id:", host.Id)
	table.AddRow("GroupName:", host.GroupName)
	table.AddRow("Name:", host.Name)
	table.AddRow("IP:", host.IP)
	table.AddRow("Processor:", host.Processor)
	table.AddRow("Memory:", humanize.Bytes(uint64(host.Memory)))
	table.AddRow("Disk:", humanize.Bytes(uint64(host.Disk)))
	table.AddRow("Platform:", host.Platform)
	table.AddRow("CreateTime:", host.CreateTime)
	table.AddRow("UpdateTime:", host.UpdateTime)
	table.AddRow("\n")
	fmt.Printf("--------------------\n")
	fmt.Printf("Host detail: \n")
	fmt.Printf("\n")
	fmt.Println(table)
}

func ListAllHost(hosts []model.Host) {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	for _, host := range hosts {
		table.AddRow("Id:", host.Id)
		table.AddRow("GroupName:", host.GroupName)
		table.AddRow("Name:", host.Name)
		table.AddRow("IP:", host.IP)
		table.AddRow("Processor:", host.Processor)
		table.AddRow("Memory:", humanize.Bytes(uint64(host.Memory)))
		table.AddRow("Disk:", humanize.Bytes(uint64(host.Disk)))
		table.AddRow("Platform:", host.Platform)
		table.AddRow("CreateTime:", host.CreateTime)
		table.AddRow("UpdateTime:", host.UpdateTime)
		table.AddRow("\n")
	}
	fmt.Printf("--------------------\n")
	fmt.Printf("Total %d hosts: \n", len(hosts))
	fmt.Printf("\n")
	fmt.Println(table)
}

func ListHost(pageResponse model.PageResponse) {
	count := pageResponse.Count
	items := pageResponse.Items
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	for _, item := range items {
		var host model.Host
		util.DecodeJSON(item, &host)
		table.AddRow("Id:", host.Id)
		table.AddRow("GroupName:", host.GroupName)
		table.AddRow("Name:", host.Name)
		table.AddRow("IP:", host.IP)
		table.AddRow("Processor:", host.Processor)
		table.AddRow("Memory:", humanize.Bytes(uint64(host.Memory)))
		table.AddRow("Disk:", humanize.Bytes(uint64(host.Disk)))
		table.AddRow("Platform:", host.Platform)
		table.AddRow("CreateTime:", host.CreateTime)
		table.AddRow("UpdateTime:", host.UpdateTime)
		table.AddRow("\n")
	}
	fmt.Printf("Total %d hosts: \n", count)
	fmt.Printf("--------------------\n")
	fmt.Println(table)
}
