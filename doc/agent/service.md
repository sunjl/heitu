```
go run main.go service get --id 1

go run main.go service listAll --groupName group01 --projectId 1 --projectName project01 --protocol http --hostName host01 --ip 192.168.1.1 --port 9000 --order createTime --direction desc

go run main.go service list --groupName group01 --projectId 1 --projectName project01 --protocol http --hostName host01 --ip 192.168.1.1 --port 9000 --page 0 --size 10 --order createTime --direction desc
```