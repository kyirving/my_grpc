package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

func FormatMsg(logleve int, message string) string {

	var dateType string

	currentTime := time.Now()

	if logleve == 1 {
		dateType = fmt.Sprintf("%s [INFO] %s", currentTime.Format("2006.01.02 15:04:05"), message)
	} else if logleve == 2 {
		dateType = fmt.Sprintf("%s [ERROR] %s", currentTime.Format("2006.01.02 15:04:05"), message)
	}
	return dateType
}

//执行命令
func Cmd(msgChan chan string, command string) {
	cmd := exec.Command("bash", "-c", command)
	var wg sync.WaitGroup

	fmt.Println("command = ", command)
	//执行命令并返回标准输出和错误输出合并的切片
	// _, err := cmd.CombinedOutput()
	// if err != nil {
	// 	fmt.Println("cmd.CombinedOutput() err = ", err)
	// }

	//返回一个在命令Start后与命令标准错误输出关联的管道。Wait方法获知命令结束后会关闭这个管道
	//一般不需要显式的关闭该管道。但是在从管道读取完全部数据之前调用Wait是错误的；同样使用StderrPipe方法时调用Run函数也是错误的
	stderrPipe, _ := cmd.StderrPipe()
	//方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，
	//一般不需要显式的关闭该管道。但是在从管道读取完全部数据之前调用Wait是错误的；同样使用StdoutPipe方法时调用Run函数也是错误的
	stdoutPipe, _ := cmd.StdoutPipe()

	//开始执行命令，但并不会等待该命令完成即返回 Wait方法会返回命令的返回状态码并在命令返回后释放相关的资源。
	if err := cmd.Start(); err != nil {
		fmt.Println("cmd.Start() err = ", err)
		log.Fatalln("cmd.Start() err = ", err)
		msgChan <- "execption"
	}

	// 标准错误输出
	wg.Add(1)
	go func() {
		defer stderrPipe.Close()
		defer wg.Done()

		scanerr := bufio.NewScanner(stderrPipe)
		for scanerr.Scan() { // 命令在执行的过程中, 实时地获取其输出
			// check out
			fmt.Println(string(scanerr.Bytes()))
			msgChan <- string(scanerr.Bytes())

		}
	}()
	//标准正确输出
	wg.Add(1)
	go func() {
		defer stdoutPipe.Close()
		defer wg.Done()

		scanout := bufio.NewScanner(stdoutPipe)
		for scanout.Scan() {
			fmt.Println(string(scanout.Bytes()))
			msgChan <- string(scanout.Bytes())

		}
	}()

	cmd.Wait()
	wg.Wait()
	//ProcessState包含一个已经存在的进程的信息，只有在调用Wait或Run后才可用
	if !cmd.ProcessState.Success() {
		msgChan <- "execption"
	}
	close(msgChan)
}

func MyCron(msgChan chan string, spec, command string) {

	c := cron.New()
	var wg sync.WaitGroup

	//@every 5s
	entryId, err := c.AddFunc(spec, func() {
		fmt.Println("command = ", command)
		//后台运行一个命令 bash -c 方式
		cmd := exec.Command("bash", "-c", command)

		stderrPipe, err := cmd.StderrPipe()
		if err != nil {
			fmt.Println("cmd.StderrPipe() err = ", err)
		}
		//方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，
		//一般不需要显式的关闭该管道。但是在从管道读取完全部数据之前调用Wait是错误的；同样使用StdoutPipe方法时调用Run函数也是错误的
		stdoutPipe, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println("cmd.stdoutPipe() err = ", err)
		}
		if err := cmd.Start(); err != nil {
			fmt.Println("cmd.Start() = ", err)
			os.Exit(0)
		}

		fmt.Println("stderrPipe = ", stderrPipe)
		fmt.Println("stdoutPipe = ", stdoutPipe)
		wg.Add(1)
		go func() {
			defer stderrPipe.Close()
			defer wg.Done()

			//创建一个流来读取管道内内容
			scanerr := bufio.NewScanner(stderrPipe)
			for scanerr.Scan() { // 命令在执行的过程中, 实时地获取其输出
				// check out
				// fmt.Println(string(scanerr.Bytes()))
				msgChan <- string(scanerr.Bytes())

			}
		}()
		//标准正确输出
		wg.Add(1)
		go func() {
			defer stdoutPipe.Close()
			defer wg.Done()

			scanout := bufio.NewScanner(stdoutPipe)
			for scanout.Scan() {
				// fmt.Println(string(scanout.Bytes()))
				msgChan <- string(scanout.Bytes())

			}
		}()
		cmd.Wait()
		//todo 需要在Wait() 方法之后调用
		if !cmd.ProcessState.Success() {
			msgChan <- "exception"
		}
		wg.Wait()
	})

	if err != nil {
		log.Fatalln("cron.AddFunc err = ", err)
	}

	fmt.Println("entryId = ", entryId)

	c.Start()
	close(msgChan)
	// 阻塞，或者使用其他延迟时间函数、
	select {}

}
