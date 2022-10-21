package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// PROJECT_NAME:Test
// DATE:2022/10/11 23:48
// USER:21005
// 驱动：go get -u github.com/go-sql-driver/mysql
// 定义一个全局对象
var db *sql.DB

// 定义一个初始化数据库函数
func initDB() (err error) {
	dsn := "root:666666@tcp(127.0.0.1:3306)/database go_db?charset=utf8mb4&parseTime=True"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用：=，我们是全局变量赋值，然后在main函数中使用全局变量dn
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否准确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// 插入数据
func insertData1() {
	sqlStr := "insert into user_tb1(username, password)values (?,?)"
	ret, err := db.Exec(sqlStr, "张三", "zs123")
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的ID
	if err != nil {
		fmt.Printf("get"+"Insert ID failed,err:%v\n", err)
		return
	}
	fmt.Printf("insert success,the id is %d\n", theID)
}

// 插入数据
func insertData2(username string, password string) {
	sqlStr := "insert into user_tb1(username, password)values (?,?)"
	ret, err := db.Exec(sqlStr, username, password)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的ID
	if err != nil {
		fmt.Printf("get"+"Insert ID failed,err:%v\n", err)
		return
	}
	fmt.Printf("insert success,the id is %d\n", theID)
}

type user struct {
	id       int
	username string
	password string
}

// 查询一条用户数据
func queryRowDemo(i int) {
	sqlStr := "select * from user_tb1 where id = ?"
	var u user
	// 确保QueryRow后调用Scan方法，否则持有的数据连接不会被释放
	err := db.QueryRow(sqlStr, i).Scan(&u.id, &u.username, &u.password)
	if err != nil {
		fmt.Printf("scan failed err,err:%v\n", err)
		return
	} else {
		fmt.Printf("id:%v,name:%s,password:%s\n", u.id, u.username, u.password)
	}
}

// 查询多条数据
func queryMultiRow() {
	sqlStr := "select id, username, password from user_tb1 where id> ?"
	rows, err := db.Query(sqlStr, 0)
	var u user
	if err != nil {
		fmt.Printf("query failed err,err:%v\n", err)
		return
	} else {
		for rows.Next() {
			err := rows.Scan(&u.id, &u.username, &u.password)
			if err != nil {
				fmt.Printf("scan failed err,err:%v\n", err)
				return
			} else {
				fmt.Printf("id:%v,username:%s,password:%s\n", u.id, u.username, u.password)
			}
		}
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()
	// 循环读取结果集中的数据

}

// 更新数据
func updateData() {
	sql := "update user_tb1 set username = ?,password = ? where id =?"
	ret, err := db.Exec(sql, "小王", "14574", 1)
	if err != nil {
		fmt.Printf("更新失败！err:%v\n", err)
		return
	}
	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("更新行失败！err:%v\n", err)
		return
	} else {
		fmt.Printf("更新成功，更新的行数为：%v\n", rows)
	}

}

// 删除数据
func delDate() {
	sql := "delete from user_tb1 where id =?"
	ret, err1 := db.Exec(sql, 1)

	if err1 != nil {
		fmt.Printf("删除行失败！err:%v\n", err1)
		return
	} else {
		rows, _ := ret.RowsAffected()
		fmt.Printf("删除成功！，删除的行数：%V\n", rows)
	}

}
func main() {
	db, err1 := sql.Open("mysql", "root:666666@/database go_db")
	if err1 != nil {
		panic(err1)
	}
	// 最大连接时长
	db.SetConnMaxLifetime(time.Minute * 3)
	// 最大连接数
	db.SetMaxOpenConns(10)
	// 空闲连接数
	db.SetMaxIdleConns(10)
	fmt.Printf("db:%T\n", db)

	err := initDB() // 调用初始化函数
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Printf("初始化成功！\n")
	}
	// 插入
	//insertData1()
	//insertData2("test", "201202541")

	// 查询
	//for i := 1; i < 5; i++ {
	//	queryRowDemo(i)
	//}
	//queryMultiRow()

	// 更新
	//updateData()

	// 删除
	//delDate()
}
