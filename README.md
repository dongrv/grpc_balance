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
##### 参数说明
```
protoc 编译器
-I 搜索路径
protobuf 当前目录下的protobuf文件夹
./protobuf/*.proto 搜索的目标文件
--go_out  输出为golang语言的文件
--go-grpc_out 除数为gRPC格式的结构文件
./protocol  生成的目标结构路径，用于golang其他程序调用
```
