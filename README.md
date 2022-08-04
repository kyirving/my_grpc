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
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux/Sserver server/main.go

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux/Sclient client/main.go
```
###### windows
```go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/windows/Sserver server/main.go

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/windows/Sclient client/main.go
```
###### mac
```go
CGO_ENABLED=0 GOOS=mac GOARCH=amd64 go build -o bin/mac/Sserver server/main.go

CGO_ENABLED=0 GOOS=mac GOARCH=amd64 go build -o bin/mac/Sclient client/main.go