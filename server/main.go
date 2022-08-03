package main

import (
	"fmt"
	"log"
	pd "my_grpc/proto/myproto"
	"my_grpc/utils"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	Address = "127.0.0.1:8080"
)

//定义cmdServer并实现约定的接口
type CmdServer struct {
}

//定义FileServer并实现约定的接口
type FileServer struct {
}

//定义消息通道，用于通讯
// var msgChan = make(chan string)
// var wg sync.WaitGroup

//ExecStream 实现Cmd服务接口，客户端待调用 响应采用流式响应，可以响应多个数据给客户端
func (c *CmdServer) ExecStream(req *pd.Request, stream pd.Cmd_ExecStreamServer) error {

	//不要定义全局的通道，否则会出现关闭的情况
	var msgChan = make(chan string)
	//检查请求方式 1 执行命令 2 执行cron
	if req.Method == "1" {
		//执行命令
		go utils.Cmd(msgChan, req.Command)

	} else if req.Method == "2" {
		//crontab
		go utils.MyCron(msgChan, req.Spec, req.Command)
	} else {
		stream.Send(&pd.Response{Message: "request method undefined"})
		return nil
	}

	for val := range msgChan {
		// 通过 send 方法不断推送数据
		err := stream.Send(&pd.Response{Message: val})
		if err != nil {
			log.Fatalf("Send error:%v", err)
			return err
		}
	}
	// 返回nil表示已经完成响应
	return nil
}

//Upload 实现File服务接口
// func (this *FileServer) UploadStream(ctx context.Context, req *pd.Request) (*pd.Response, error) {
// 	resp := new(pd.Response)
// 	resp.Message = fmt.Sprintf("cmd: %s.", req.Command)

// 	return resp, nil
// }

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	// 实例化grpc Server
	grpcServer := grpc.NewServer()

	// 注册CmdService
	pd.RegisterCmdServer(grpcServer, &CmdServer{})

	//注册FileServer
	// pd.RegisterFileServer(grpcServer, &FileServer{})

	fmt.Println("服务端开始监听～～～")
	log.Printf("[INFO] my-grpc Server Start at Pid:%d Address:%s\n", os.Getpid(), Address)
	// grpclog.Println("Listen on " + Address)
	grpcServer.Serve(listen)

}
