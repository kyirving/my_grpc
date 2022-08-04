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

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	otype byte
)

func main() {
	fmt.Println("------------ 请 选 择 模 式 -----------")
	fmt.Println("             1 : 命 令 执 行")
	fmt.Println("             2 : 文 件 下 发")

	fmt.Scanln(&otype)

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
	fmt.Scanln(&address)

	fmt.Println("请输入命令:")
	fmt.Scanln(&commmand)
	if method == "1" {

	} else if method == "2" {
		fmt.Println("请输入执行规则:")
		fmt.Scanln(&spec)
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

	var (
		address string
	)

	//连接资源服务
	SouceConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer SouceConn.Close()
	err = DownloadFile(SouceConn)
	if err != nil {
		log.Println("DownloadFile err = ", err)
		return
	}

	//连接上传的服务
	ToConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer ToConn.Close()

	err = UploadFile(ToConn)
	if err != nil {
		log.Println("UploadFile err = ", err)
		return
	}

}

func DownloadFile(conn *grpc.ClientConn) error {
	//1 先下载
	client := pd.NewFileClient(conn)
	req := &pd.DlRequest{
		Filepath: "/Users/wuh/study/go/src/my_grpc/server/test.php",
	}

	stream, err := client.DownloadFile(context.Background(), req)
	if err != nil {
		log.Fatalf("could not echo: %v", err)
		return errors.New(error.Error(err))
	}

	// for循环获取服务端推送的消息
	file, err := os.OpenFile("../script/test.php", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("os.OpenFile() err = ", err)
		return errors.New(error.Error(err))
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

		n, err := write.Write(resp.Data)
		fmt.Println("n = ", n)
		fmt.Println("err = ", err)

		log.Printf("Recv data:%v", resp)
	}
	write.Flush()

	return nil
}

func UploadFile(conn *grpc.ClientConn) error {
	client := pd.NewFileClient(conn)

	filepath := "../script/test.php"
	filename := "upload.php"

	//得到流的句柄
	stream, err := client.UploadFile(context.Background())
	if err != nil {
		fmt.Println("os.OpenFile() err = ", err)
		return errors.New(error.Error(err))
	}

	if !utils.FileExists(filepath) {
		return errors.New("文件不存在")
	}

	file, err := os.Open(filepath)
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

		fmt.Println("client Send Upload Data = ", string(buf[:n]))
		stream.Send(&pd.UpRequest{
			Data:     buf[:n],
			Filename: filename,
		})
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println("stream.CloseAndRecv() err =", err)
	}

	fmt.Println("resp = ", resp)
	return nil
}
