package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done()
	fmt.Printf("hello %d\n", i)
}

func a()  {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("a", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("b", i)
	}
}

func run(){
	runtime.GOMAXPROCS(2) //默认cpu的逻辑核心数，默认跑满cpu
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}

func receiver(c chan int) {
	r := <- c
	fmt.Println("从通道接收到：", r)
}

func emitter() {
	c := make(chan int)
	go receiver(c)
	c <- 10
	fmt.Print("发送完成")
}

func writer(out chan<- int) { //单向通道,只能往内写
	for i := 0; i < 10; i++{
		out <- i
	}
	fmt.Println("已经发送")
	close(out)
}

func reader(out chan<- int, in <-chan int) {
	for i := range in{  //只能读取
		out <- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in{
		fmt.Print(i)
	}
}

func main() {
	var wg sync.WaitGroup //等待组不需要初始化

	var i int = -1
	var file string

	// 在命令行上获取一个文件列表
	for i, file = range os.Args[1:]{
		wg.Add(1)  //告诉等待组，正在等待多个压缩操作
		go func(file string) {
			compress(file)
			wg.Done()  // 通知等待组压缩完成
		}(file)
	}
	wg.Wait() //直到所有的goroutine完成，主程序开始执行
	fmt.Printf("共压缩了%d files", i+1)
}

func compress(filename string) error{
	// 读取打开的一个源文件
	in, err := os.Open(filename)
	if err != nil{
		return err
	}
	defer in.Close()
	// 打开一个目标文件，并将.gz扩展名添加到文件上
	out, err := os.Create(filename + ".gz")
	if err != nil{
		return err
	}
	defer out.Close()
	// 压缩工具将数据压缩成包并写入底层文件
	gout := gzip.NewWriter(out)
	// 进行所有的复制操作
	_, err = io.Copy(gout, in)
	gout.Close()

	return err
}



