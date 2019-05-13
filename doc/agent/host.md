```
go run main.go host add --groupName group01 --ethernet eth0 --disk sda1

go run main.go host get --name host01

go run main.go host listAll --groupName group01 --order createTime --direction desc

go run main.go host list --groupName group01 --page 0 --size 10 --order createTime --direction desc

go run main.go host update --groupName group01 --ethernet eth0 --disk sda1

go run main.go host delete --name host01

go run main.go host register
```