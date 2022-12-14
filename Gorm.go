package main

import (
	"gorm.io/gorm"
)

// PROJECT_NAME:Test
// DATE:2022/10/18 23:37
// USER:21005

// go get -u gorm.io/gorm
// go get -u gorm.io/driver/sqlite
// go get -u gorm.io/driver/mysql
type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	//dsn := "root:666666@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}

	// 迁移 schema
	//db.AutoMigrate(&Product{})

	//Create
	//db.Create(&Product{Code: "D42", Price: 100})

	// Read
	//var product Product
	//db.First(&product, 1) // 根据整型主键查找
	//fmt.Printf("Product %v\n", product)
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	//fmt.Printf("Product %v\n", product)

	//// Update - 将 product 的 price 更新为 200
	//db.First(&product, 1) // 根据整型主键查找
	//db.Model(&product).Update("price", 200)

	//// Update - 更新多个字段
	//db.First(&product, 1) // 根据整型主键查找
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F44"})

	// Delete - 删除 product
	//db.First(&product, 1) // 根据整型主键查找
	//db.Delete(&product, 1)
}
