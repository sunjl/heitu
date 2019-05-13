```
go run main.go config get --id 1

go run main.go config listAll --projectId 1 --projectName project01 --version 1.0 --environment dev --order createTime --direction desc

go run main.go config list --projectId 1 --projectName project01 --version 1.0 --environment dev --page 0 --size 10 --order createTime --direction desc
```