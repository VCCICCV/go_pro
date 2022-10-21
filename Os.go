package main

import (
	"fmt"
	"os"
)

// PROJECT_NAME:Test
// DATE:2022/10/6 23:48
// USER:21005
// 创建文件
func createTestProject() {
	f, err := os.Create("project.txt")
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Printf("createTestProject\nf.name:%v\n", f.Name())
	}
	err2 := os.MkdirAll("c/v/f\n", os.ModePerm)
	if err2 != nil {
		fmt.Printf("yes!!!:%v\n", err2)
	}
}

// 创建目录
func makeDir() {
	err := os.Mkdir("project", os.ModePerm)
	if err != nil {
		fmt.Printf("f.name:%v\n", err)
	} else {
		fmt.Printf("makeDir yes!!!\n")
	}
}

// 删除目录或者文件
func remove() {
	err := os.Remove("project.txt")
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Printf("remove yes!!!\n")
	}
}

// 获得工作目录
func getWd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Printf("wd:%v\n", dir)
	}
}

// 修改工作目录

func chWd() {
	err := os.Chdir("d:/")
	dir, _ := os.Getwd()
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Printf("修改成功！")
		fmt.Printf("dir:%v\n", dir)
	}
}

// 获取临时目录
func getTemp() {
	s := os.TempDir()
	fmt.Printf("%s\n", s)
}

// 重命名
func rename() {
	err := os.Rename("project.txt", "Rename.txt")
	if err != nil {
		fmt.Printf("出错了:%s\n", err)
	} else {
		fmt.Println("修改成功!")
	}
}

// 读文件
func readFile() {
	r, err := os.ReadFile("main.go")
	if err != nil {
		fmt.Printf("出错了：%s", err)
	} else {
		fmt.Printf("r:\n%v\n", string(r[:]))
	}
}

// 写文件
func writeFile() {
	s := "努力学习，发展乡村！"
	os.WriteFile("Rename.txt", []byte(s), os.ModePerm)
}
func main() {
	//createTestProject()
	//makeDir()
	//remove()
	//getWd()
	////chWd()
	//getTemp()
	//rename()
	readFile()
	writeFile()
}
