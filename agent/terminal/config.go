package terminal

import (
	"fmt"
	"github.com/gosuri/uitable"
	"github.com/jishufan/heitu/client/model"
	"github.com/jishufan/heitu/common/util"
)

func ShowConfig(config model.Config) {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	table.AddRow("Id:", config.Id)
	table.AddRow("ProjectId:", config.ProjectId)
	table.AddRow("ProjectName:", config.ProjectName)
	table.AddRow("Version:", config.Version)
	table.AddRow("Environment:", config.Environment)
	table.AddRow("FileName:", config.FileName)
	table.AddRow("CreateTime:", config.CreateTime)
	table.AddRow("UpdateTime:", config.UpdateTime)
	table.AddRow("\n")
	fmt.Printf("--------------------\n")
	fmt.Printf("Config detail: \n")
	fmt.Printf("\n")
	fmt.Println(table)
}

func ListAllConfig(configs []model.Config) {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	for _, config := range configs {
		table.AddRow("Id:", config.Id)
		table.AddRow("ProjectId:", config.ProjectId)
		table.AddRow("ProjectName:", config.ProjectName)
		table.AddRow("Version:", config.Version)
		table.AddRow("Environment:", config.Environment)
		table.AddRow("FileName:", config.FileName)
		table.AddRow("CreateTime:", config.CreateTime)
		table.AddRow("UpdateTime:", config.UpdateTime)
		table.AddRow("\n")
	}
	fmt.Printf("--------------------\n")
	fmt.Printf("Total %d configs: \n", len(configs))
	fmt.Printf("\n")
	fmt.Println(table)
}

func ListConfig(pageResponse model.PageResponse) {
	count := pageResponse.Count
	items := pageResponse.Items
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	for _, item := range items {
		var config model.Config
		util.DecodeJSON(item, &config)
		table.AddRow("Id:", config.Id)
		table.AddRow("ProjectId:", config.ProjectId)
		table.AddRow("ProjectName:", config.ProjectName)
		table.AddRow("Version:", config.Version)
		table.AddRow("Environment:", config.Environment)
		table.AddRow("FileName:", config.FileName)
		table.AddRow("CreateTime:", config.CreateTime)
		table.AddRow("UpdateTime:", config.UpdateTime)
		table.AddRow("\n")
	}
	fmt.Printf("Total %d configs: \n", count)
	fmt.Printf("--------------------\n")
	fmt.Println(table)
}
