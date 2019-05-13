```
go get -u encoding/json
go get -u github.com/coreos/etcd/clientv3
go get -u github.com/dustin/go-humanize
go get -u github.com/go-sql-driver/mysql
go get -u github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/gosuri/uitable
go get -u github.com/jinzhu/gorm
go get -u github.com/labstack/echo
go get -u github.com/Masterminds/glide
go get -u github.com/mitchellh/mapstructure
go get -u github.com/parnurzeal/gorequest
go get -u github.com/rs/cors
go get -u github.com/satori/go.uuid
go get -u github.com/spf13/cobra/cobra
go get -u github.com/spf13/viper
go get -u github.com/tylerb/graceful
go get -u golang.org/x/crypto/bcrypt
go get -u golang.org/x/sys/unix
go get -u google.golang.org/grpc

glide get encoding/json
glide get github.com/coreos/etcd/clientv3
glide get github.com/dustin/go-humanize
glide get github.com/go-sql-driver/mysql
glide get github.com/golang/protobuf/proto
glide get github.com/golang/protobuf/protoc-gen-go
glide get github.com/gosuri/uitable
glide get github.com/jinzhu/gorm
glide get github.com/labstack/echo
glide get github.com/mitchellh/mapstructure
glide get github.com/parnurzeal/gorequest
glide get github.com/rs/cors
glide get github.com/satori/go.uuid
glide get github.com/spf13/cobra/cobra
glide get github.com/spf13/viper
glide get github.com/tylerb/graceful
glide get golang.org/x/crypto/bcrypt
glide get golang.org/x/sys/unix
glide get google.golang.org/grpc

glide install
glide update

for path in agent client common server; do
    find $path -type f -name "*.go" -exec go fmt {} \;
done

cd server && go run main.go
```