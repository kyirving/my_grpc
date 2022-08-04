package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	pd "my_grpc/proto/myproto"
	"my_grpc/utils"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	otype  byte
	reader *bufio.Reader
)

func main() {
	fmt.Println("------------ 请 选 择 模 式 -----------")
	fmt.Println("             1 : 命 令 执 行")
	fmt.Println("             2 : 文 件 下 发")

	fmt.Scanln(&otype)
	reader = bufio.NewReader(os.Stdin)
	switch otype {
	case 1:
		Command()
	case 2:
		FileSource()
		// File(conn)
	}
}

//执行命令方法
func Command() {
	var (
		method   string
		commmand string
		spec     string
		address  string
	)

	req := &pd.Request{}
	fmt.Println("------------ 请输入执行对象 ------------")
	fmt.Println("             1 普 通 命 令 ")
	fmt.Println("             2 定 时 任 务 ")
	fmt.Scanln(&method)

	fmt.Println("请输入远程主机地址及端口:")
	result, _, _ := reader.ReadLine()
	address = string(result)
	if address == "" {
		log.Fatalln("远程主机地址不能为空")
	}

	fmt.Println("请输入命令:")
	result, _, _ = reader.ReadLine()
	commmand = string(result)
	if method == "1" {

	} else if method == "2" {
		fmt.Println("请输入执行规则:")
		result, _, _ = reader.ReadLine()
		spec = string(result)
	} else {
		fmt.Println("method undefined!!!")
		return
	}
	req.Method = method
	req.Command = commmand
	req.Spec = spec

	// 客户端连接gRPC服务地址
	log.Printf("my-grpc Client grpc.Dial at Address:%s\n", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	// 初始化客户端
	client := pd.NewCmdClient(conn)

	//调用方法 获取stream
	stream, err := client.ExecStream(context.Background(), req)
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}

	// for循环获取服务端推送的消息
	for {
		// 通过 Recv() 不断获取服务端send()推送的消息
		resp, err := stream.Recv()
		// err==io.EOF则表示服务端关闭stream了 退出
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Client Recv error:%v", err)
			continue
		}
		log.Printf("Client Recv data:%v\n", resp.GetMessage())
	}
}

func FileSource() {

	fmt.Println("请输入数据源主机地址及端口:")
	result, _, _ := reader.ReadLine()
	address := string(result)

	if address == "" {
		log.Fatalln("数据源主机地址不能为空")
	}

	//连接资源服务
	SouceConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer SouceConn.Close()
	localPath, err := DownloadFile(SouceConn)
	if err != nil {
		log.Fatalln("DownloadFile err = ", err)
		return
	}

	fmt.Println("请输入分发主机地址及端口:")
	result, _, _ = reader.ReadLine()
	toAddress := string(result)
	if address == "" {
		log.Fatalln("分发主机地址不能为空")
	}

	//连接上传的服务
	ToConn, err := grpc.Dial(toAddress, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer ToConn.Close()

	err = UploadFile(ToConn, localPath)
	if err != nil {
		log.Println("UploadFile err = ", err)
		return
	}

}

func DownloadFile(conn *grpc.ClientConn) (string, error) {
	//1 先下载

	fmt.Println("请输入数据源主机文件路径:")
	result, _, _ := reader.ReadLine()
	Filepath := string(result)
	if Filepath == "" {
		return "", errors.New("数据源主机文件路不能为空")
	}

	client := pd.NewFileClient(conn)
	req := &pd.DlRequest{
		Filepath: Filepath,
	}

	stream, err := client.DownloadFile(context.Background(), req)
	if err != nil {
		log.Fatalf("could not echo: %v", err)
		return "", errors.New(error.Error(err))
	}

	// for循环获取服务端推送的消息
	localFilePath := "../script/" + filepath.Base(Filepath)
	file, err := os.OpenFile(localFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("os.OpenFile() err = ", err)
		return "", errors.New(error.Error(err))
	}
	write := bufio.NewWriter(file)
	for {
		// 通过 Recv() 不断获取服务端send()推送的消息
		resp, err := stream.Recv()
		// err==io.EOF则表示服务端关闭stream了 退出
		if err == io.EOF {
			log.Println("server closed")
			break
		}
		if err != nil {
			log.Printf("Recv error:%v", err)
			continue
		}

		//服务端响应失败
		if resp.Code != utils.RESP_SUCC {
			return "", errors.New(resp.Msg)
		}

		write.Write(resp.Data)
		log.Printf("Recv data:%v", resp)
	}
	write.Flush()
	return localFilePath, nil
}

func UploadFile(conn *grpc.ClientConn, localPath string) error {
	client := pd.NewFileClient(conn)

	fmt.Println("请输入目标主机文件路径:")
	result, _, _ := reader.ReadLine()
	RemoteFilePath := string(result)
	if RemoteFilePath == "" {
		return errors.New("目标主机文件路径不能为空")
	}
	// RemoteFilePath := "/tmp/upload.php"

	//得到流的句柄
	stream, err := client.UploadFile(context.Background())
	if err != nil {
		fmt.Println("os.OpenFile() err = ", err)
		return errors.New(error.Error(err))
	}

	if !utils.FileExists(localPath) {
		return errors.New("文件不存在")
	}

	file, err := os.Open(localPath)
	if err != nil {
		return errors.New("打开文件失败")
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
		stream.Send(&pd.UpRequest{
			Data:     buf[:n],
			Filepath: RemoteFilePath,
		})
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println("stream.CloseAndRecv() err =", err)
	}

	if resp.Code == utils.RESP_SUCC {
		log.Println("文件下发成功")
	} else {
		log.Println("文件下发失败:" + resp.Msg)
	}
	return nil
}
