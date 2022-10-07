### proto cmd
##### 安装`gRPC`插件
```
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```
##### 生成`go-grpc`
```
protoc -I protobuf  ./protobuf/*.proto  --go_out=./protocol
protoc -I protobuf  ./protobuf/*.proto  --go-grpc_out=./protocol
```