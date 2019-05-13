```
cd common/protobuf
protoc -I=$GOPATH/src -I. --go_out=plugins=grpc:. *.proto
```