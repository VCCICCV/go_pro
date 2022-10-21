package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// PROJECT_NAME:Test
// DATE:2022/10/9 1:11
// USER:21005
func testCopy() {
	r := strings.NewReader("hello world")
	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
}
func TestCopyBuf() {
	r1 := strings.NewReader("first reader!\n")
	r2 := strings.NewReader("second reader!\n")
	buf := make([]byte, 8)
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}
	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}
}

// 多个reader
func testReader() {
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("first reader\n")
	r3 := strings.NewReader("first reader\n")
	r4 := strings.NewReader("first reader\n")
	r := io.MultiReader(r1, r2, r3, r4)

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err) //直接退出不执行defer
		defer fmt.Printf("test")
	}
}
func testBufIo() {
	s := strings.NewReader("ASDFGHJKL")
	br := bufio.NewReader(s)
	c, _ := br.ReadByte()
	fmt.Printf("%c\n", c)
	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)
	br.UnreadByte() // 吐出来一个
	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)
	c, _ = br.ReadByte()
	fmt.Printf("%c\n", c)
}
func BufReaderWriter() {
	b := bytes.NewBuffer(make([]byte, 0))
	s := strings.NewReader("ASDFGH")

	bw := bufio.NewWriter(b)
	br := bufio.NewReader(s)
	rw := bufio.NewReadWriter(br, bw)
	p, _ := rw.ReadString('\n')
	fmt.Println(string(p))
	rw.WriteString("asdf")
	rw.Flush()
	fmt.Println(b)

}
func Scanner() {
	s := strings.NewReader("AS DF GH")
	bs := bufio.NewScanner(s)
	bs.Split(bufio.ScanWords) // 以空格进行分割
	for bs.Scan() {
		fmt.Println(bs.Text())
	}
}
func main() {
	//r := strings.NewReader("hello word")
	//buf := make([]byte, 20)
	//r.Read(buf)
	//fmt.Printf("string(buf):%v\n", string(buf))

	//testCopy().
	//TestCopyBuf()
	//testReader()
	//testBufIo()
	//BufReaderWriter()
	Scanner()
}
