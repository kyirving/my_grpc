package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	pd "my_grpc/proto/myproto"
	"my_grpc/utils"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	Address = ":8080"
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

//实现FileServer服务的DownloadFile方法
func (f *FileServer) DownloadFile(req *pd.DlRequest, stream pd.File_DownloadFileServer) error {

	fmt.Println("filepath = ", req.Filepath)
	//先检测文件是否存在
	if !utils.FileExists(req.Filepath) {
		stream.Send(&pd.DlResponse{
			Code: utils.RESP_NOT_FOUND,
			Msg:  "文件或目录不存在",
		})
		return nil
	}

	file, err := os.Open(req.Filepath)
	if err != nil {
		stream.Send(&pd.DlResponse{
			Code: utils.RESP_SYSTEM_BUSY,
			Msg:  "打开文件错误",
		})
		return nil
	}
	//创建缓存区，用于读取文件内容
	buf := make([]byte, 10)
	for {
		n, err := file.Read(buf)
		//读取完毕
		if err == io.EOF {
			log.Println("file read Done")
			break
		}

		if err != nil {
			log.Println("file read err = ", err)
			break
		}
		stream.Send(&pd.DlResponse{
			Data: buf[:n],
			Code: utils.RESP_SUCC,
		})
	}
	return nil
}

func (f *FileServer) UploadFile(stream pd.File_UploadFileServer) error {

	//读取客户端上传的内容
	var file *os.File
	var filepath string
	var write *bufio.Writer
	defer file.Close()
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("server stream read Done")
			goto END
		}
		if err != nil {
			return stream.SendAndClose(&pd.UpResponse{
				Code: utils.RESP_SYSTEM_BUSY,
				Msg:  "server read err :" + error.Error(err),
			})

			log.Println("stream read err = ", err)
			break
		}
		//打开文件
		if file == nil {
			filepath = req.Filepath
			file, err = os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				log.Println("Server OpenFile err = ", err)
			}
		}

		write = bufio.NewWriter(file)
		write.Write(req.Data)
		write.Flush() //todo 先每循环一次执行下吧
	}
END:
	write.Flush()
	return stream.SendAndClose(&pd.UpResponse{
		Filepath: filepath,
		Code:     utils.RESP_SUCC,
	})
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
	pd.RegisterFileServer(grpcServer, &FileServer{})

	fmt.Println("服务端开始监听～～～")
	log.Printf("[INFO] my-grpc Server Start at Pid:%d Address:%s\n", os.Getpid(), Address)
	// grpclog.Println("Listen on " + Address)
	grpcServer.Serve(listen)

}
