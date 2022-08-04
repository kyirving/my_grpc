# my_grpc

#### 安装说明
编译Linux,Windows和Mac环境下可执行程序

```go
git clone https://github.com/kyirving/my_grpc.git
cd my_grpc
go mod tidy 
```

###### linux
```go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux/myServer server/main.go

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux/myClient client/main.go
```
###### windows
```go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/windows/myServer server/main.go

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/windows/myClient client/main.go
```
###### mac
```go
CGO_ENABLED=0 GOOS=mac GOARCH=amd64 go build -o bin/mac/myServer server/main.go

CGO_ENABLED=0 GOOS=mac GOARCH=amd64 go build -o bin/mac/myClient client/main.go