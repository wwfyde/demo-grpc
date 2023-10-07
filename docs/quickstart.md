# 快速开始
- [grpc-go-quickstart](https://grpc.io/docs/languages/go/quickstart/)
- [Go Generated Code Guide](https://protobuf.dev/reference/go/go-generated/)

## 插件安装

```shell
# install protocol buffers compiler
brew install protobuf
```

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
## 更新PATH
```shell
export PATH="$PATH:$(go env GOPATH)/bin"
```

## 生成grpc code

```shell
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto

# If the paths=source_relative flag is specified, the output file is placed in the same relative directory as the input file. For example, an input file protos/buzz.proto results in an output file at protos/buzz.pb.go.
```

## 开发思路

### Server
- 创建服务网络监听器Listener
- 创建grpc服务
- 注册service到grpc.server
- 启动服务Serve函数,用于处理从Listener发起的请求, 并为每个请求创建一个goroutine来处理该请求 
