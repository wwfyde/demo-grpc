# 快速开始
- [grpc-go-quickstart](https://grpc.io/docs/languages/go/quickstart/)
## 插件安装

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
```