package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// PROJECT_NAME:Test
// DATE:2022/10/20 19:35
// USER:21005
var db1 *gorm.DB

func init() {
	dsn := "root:666666@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	// 全局配置
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{DryRun: true})
	// DryRun: true只生成sql语句不执行
	if err != nil {
		panic(err)
	}
	db1 = d
}
func test1() {
	// 会话配置
	db1.Session(&gorm.Session{DryRun: true})
}
func main() {
	test1()
}
