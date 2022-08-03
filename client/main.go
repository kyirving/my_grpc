package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	pd "my_grpc/proto/myproto"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:8080"
)

var (
	otype byte
)

func main() {
	// 客户端连接gRPC服务地址
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	fmt.Println("------------ 请 选 择 模 式 -----------")
	fmt.Println("             1 : 命 令 执 行")
	fmt.Println("             2 : 文 件 下 发")

	fmt.Scanln(&otype)

	switch otype {
	case 1:
		Command(conn)
	case 2:
		File(conn)
	}
}

//执行命令方法
func Command(conn *grpc.ClientConn) {
	// 初始化客户端
	client := pd.NewCmdClient(conn)
	// 调用方法
	req := &pd.Request{
		Method:  "2", //1 是命令 2是cron
		Command: "echo 222 >> ./tast2.log",
		Spec:    "*/1 * * * *",
	}
	//调用获取stream
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
			log.Println("server closed")
			break
		}
		if err != nil {
			log.Printf("Recv error:%v", err)
			continue
		}
		log.Printf("Recv data:%v", resp.GetMessage())
	}
}

func File(conn *grpc.ClientConn) {
	//1 先下载
	client := pd.NewFileClient(conn)
	req := &pd.DlRequest{
		Filepath: "/Users/wuh/study/go/src/my_grpc/server/test.php",
	}

	stream, err := client.DownloadFile(context.Background(), req)
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}

	// for循环获取服务端推送的消息
	file, err := os.OpenFile("../script/test.php", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("os.OpenFile() err = ", err)
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
}
