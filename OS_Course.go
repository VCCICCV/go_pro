package main

import (
	"fmt"
	"os"
	"time"
)

// PROJECT_NAME:Test
// DATE:2022/10/9 0:18
// USER:21005
func main() {
	// 获得当前正在运行的进程
	fmt.Printf("os.Getpid:%v\n", os.Getegid())
	// 父id
	fmt.Printf("os.Getppid:%v\n", os.Getppid())
	// 设置新进程的属性
	attr := &os.ProcAttr{
		// files指定的新进程继承活动文件对象
		// 前三个分别为，标准输入，标准输出，标准错误输出
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		// 新进程的环境变量
		Env: os.Environ(),
	}
	// 开始一个新进程
	p, err1 := os.StartProcess("D:\\小工具\\mpv\\mpv.exe", []string{"D:\\小工具\\mpv\\mpv.exe", "D:\\a.txt"}, attr)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(p)
	fmt.Println("进程ID：", p.Pid)
	// 通过id查找进程
	p2, _ := os.FindProcess(p.Pid)
	fmt.Println(p2)
	// 等待10秒，执行函数
	time.AfterFunc(time.Second*10, func() {
		// 向P进程发送退出信号
		p.Signal(os.Kill)
	})
	// 等待进程p的退出，返回进程状态
	ps,_ := p.Wait()
	fmt.Println(ps.String())
}
