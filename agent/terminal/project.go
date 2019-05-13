package terminal

import (
	"fmt"
	"github.com/gosuri/uitable"
	"github.com/jishufan/heitu/client/model"
	"github.com/jishufan/heitu/common/util"
)

func ShowProject(project model.Project) {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	table.AddRow("Id:", project.Id)
	table.AddRow("GroupName:", project.GroupName)
	table.AddRow("Name:", project.Name)
	table.AddRow("Git:", project.Git)
	table.AddRow("CreateTime:", project.CreateTime)
	table.AddRow("UpdateTime:", project.UpdateTime)
	table.AddRow("\n")
	fmt.Printf("--------------------\n")
	fmt.Printf("Project detail: \n")
	fmt.Printf("\n")
	fmt.Println(table)
}

func ListAllProject(projects []model.Project) {
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	for _, project := range projects {
		table.AddRow("Id:", project.Id)
		table.AddRow("GroupName:", project.GroupName)
		table.AddRow("Name:", project.Name)
		table.AddRow("Git:", project.Git)
		table.AddRow("CreateTime:", project.CreateTime)
		table.AddRow("UpdateTime:", project.UpdateTime)
		table.AddRow("\n")
	}
	fmt.Printf("--------------------\n")
	fmt.Printf("Total %d projects: \n", len(projects))
	fmt.Printf("\n")
	fmt.Println(table)
}

func ListProject(pageResponse model.PageResponse) {
	count := pageResponse.Count
	items := pageResponse.Items
	table := uitable.New()
	table.MaxColWidth = 60
	table.Wrap = true
	for _, item := range items {
		var project model.Project
		util.DecodeJSON(item, &project)
		table.AddRow("Id:", project.Id)
		table.AddRow("GroupName:", project.GroupName)
		table.AddRow("Name:", project.Name)
		table.AddRow("Git:", project.Git)
		table.AddRow("CreateTime:", project.CreateTime)
		table.AddRow("UpdateTime:", project.UpdateTime)
		table.AddRow("\n")
	}
	fmt.Printf("Total %d projects: \n", count)
	fmt.Printf("--------------------\n")
	fmt.Println(table)
}
