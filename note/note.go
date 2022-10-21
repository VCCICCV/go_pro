package note

import "fmt"

func Test() {
	fmt.Print("this is test!")

}
func IfElse() {
	var age uint8
	fmt.Print("请输入您的年龄：")
	fmt.Scanln(&age)
	if age < 13 {
		fmt.Print("小朋友别卷了")
	} else {
		fmt.Println("大学生一定要仔细学！")
	}
}
func SwitchCase() {
	var age uint8
	fmt.Print("请输入您的年龄：")
	fmt.Scanln(&age)
	switch {
	case age < 13:
		fmt.Print("小朋友别学")
		fallthrough // 继续匹配
	case age < 25:
		fmt.Print("25以内的大学生")
	default:
		break
	}
}
func Swap(x, y string) (string, string) {
	return x, y
}
func Array() {
	var balance = [5]float64{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Printf("array=%v", balance)
}
func Pointer() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)
	fmt.Printf("ip 变量储存的指针地址: %X\n", ip)
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	// nil
	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr)
}

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func Type_struct() {
	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
}
func Test_2() {
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", Book1.title)
	fmt.Printf("Book 1 author : %s\n", Book1.author)
	fmt.Printf("Book 1 subject : %s\n", Book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", Book1.book_id)

	/* 打印 Book2 信息 */
	fmt.Printf("Book 2 title : %s\n", Book2.title)
	fmt.Printf("Book 2 author : %s\n", Book2.author)
	fmt.Printf("Book 2 subject : %s\n", Book2.subject)
	fmt.Printf("Book 2 book_id : %d\n", Book2.book_id)

	var Book3 Books /* 声明 Book1 为 Books 类型 */
	var Book4 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book3.title = "Go 语言"
	Book3.author = "www.runoob.com"
	Book3.subject = "Go 语言教程"
	Book3.book_id = 6495407

	/* book 2 描述 */
	Book4.title = "Python 教程"
	Book4.author = "www.runoob.com"
	Book4.subject = "Python 语言教程"
	Book4.book_id = 6495700

	/* 打印 Book1 信息 */
	printBook(&Book3)

	/* 打印 Book2 信息 */
	printBook(&Book4)
}
func printBook(book *Books) {
	fmt.Printf("这本书:%s\n", book)
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}
func Array_2() {
	var two_Demo [3][4]int = [3][4]int{
		{12, 2, 3, 5},
		{4, 5, 6, 5},
		{8, 9, 4},
	}
	for i, v := range two_Demo {
		for i2, v2 := range v {
			fmt.Printf("a[%v][%v]=%v\t", i, i2, v2)
		}
		fmt.Println()
	}
} // 接口
type textMes struct {
	Type string
	text string
}

func (m *textMes) setText() {
	m.text = "hello world"
}

type imgMes struct {
	Type string
	Img  string
}

func (im *imgMes) setImg() {
	im.Img = "images"

}

type Mes interface {
	setType()
}

func (tm *textMes) setType() {
	tm.Type = "文字消息"
}
func (im *imgMes) setType() {
	im.Type = "图片消息"
}
func SendMes(m Mes) {
	m.setType()
	fmt.Println("m=", m)
}
func Interface() {
	tm := textMes{}
	SendMes(&tm)
	im := imgMes{}
	SendMes(&im)

}
