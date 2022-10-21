package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// PROJECT_NAME:Test
// DATE:2022/10/19 11:26
// USER:21005
var dbm *gorm.DB

type User struct {
	// 要大写才能创建进去
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

func init() {
	dsn := "root:666666@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		dbm = d
		fmt.Printf("connect success! d:%v\n", d)
	}
}
func createTable() {
	dbm.AutoMigrate(&User{})

}
func insert1() {
	user := User{
		Name:     "tom",
		Age:      21,
		Birthday: time.Now(),
	}
	result := dbm.Create(&user)
	fmt.Printf("result.RowsAffected:%v\n", result.RowsAffected)
	fmt.Printf("user.ID:%v\n", user.ID)
}
func insert2() {
	user := User{
		Name:     "jack",
		Age:      23,
		Birthday: time.Now(),
	}
	// 选择添加
	dbm.Select("Name", "Age", "CreateAt").Create(&user)

}
func insert3() {
	var users = []User{{Name: "jack1", Age: 23}, {Name: "jack2", Age: 43}, {Name: "jack3", Age: 24}}
	dbm.Create(&users)
}
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Printf("BeforeCreate******创建之前****\n")
	return
}

func insert4() {
	dbm.Model(&User{}).Create(map[string]interface{}{
		"Name": "小王", "Age": 27,
	})
}
func select1() {
	var user *User
	tx := dbm.First(&user)
	fmt.Printf("tx:%T\n", tx.RowsAffected)
	fmt.Printf("SELECT user.ID:%v\n", user.ID)
	fmt.Printf("SELECT user:%v\n", user.Name)

	//dbm.Take(&user)
	//fmt.Printf("SELECT user.ID:%v\n", user.ID)
	//fmt.Printf("SELECT user:%v\n", user.Name)

	//dbm.Last(&user)
	//fmt.Printf("SELECT user.ID:%v\n", user.ID)
	//fmt.Printf("SELECT user:%v\n", user.Name)

}
func select2() {
	//var user *User
	//dbm.First(&user, 5)
	//fmt.Printf("SELECT user:%v\n", user.ID)
	//dbm.First(&user, "10")
	//fmt.Printf("SELECT user:%v\n", user.ID)

	var users []User
	dbm.Find(&users, []int{1, 2, 3})
	for _, users := range users {
		fmt.Printf("SELECT user:%v\n", users.ID)
	}
}
func update1() {
	var user User
	dbm.First(&user)
	user.Age = 54
	user.Name = "Updated"
	dbm.Save(&user)
}
func update2() {
	dbm.Model(&User{}).Where("age = ?", 24).Update("name", "hello")
}
func main() {
	//createTable()

	//insert1()
	//insert2()
	//insert3()
	//insert4()

	//select1()
	//select2()

	//update1()
	update2()

}
