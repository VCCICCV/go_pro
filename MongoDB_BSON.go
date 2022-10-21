package main

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
)

// PROJECT_NAME:Test
// DATE:2022/10/17 11:15
// USER:21005
var client *mongo.Client

func initDb2() {
	//// 设置客户端连接配置
	//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	//co := options.Client().ApplyURI("mongodb://localhost:27017")
	//mongo.Connect(context.TODO(), co)
	//// 连接到MongoDB
	//var err error
	//client, err := mongo.Connect(context.TODO(), clientOptions)
	//if err != nil {
	//	log.Fatal(err)
	//} else {
	//	fmt.Printf("client:%T\n", client)
	//}
	//// 检查连接
	//err = client.Ping(context.TODO(), nil)
	//if err != nil {
	//	log.Fatal(err)
	//} else {
	//	fmt.Println("Connected to MongoDB success!")
	//}

	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("连接成功！\n")
	}
	client = c
}

type Student struct {
	Name string
	Age  int32
}

func insert() {
	initDb2()
	s := Student{Name: "test", Age: 21}

	collection := client.Database("golang_db").Collection("Student")
	//defer collection.Close()
	insertResult, err := collection.InsertOne(context.TODO(), s)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Inserted a single document:%v\n", insertResult.InsertedID)
	}

}
func InsertMoney() {
	initDb2()
	c := client.Database("golang_db").Collection("Student")
	s1 := Student{Name: "test1", Age: 21}
	s2 := Student{Name: "test2", Age: 24}
	stus := []interface{}{s1, s2}
	imr, err := c.InsertMany(context.TODO(), stus)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Printf("imr.InsertedIDS:%v\n", imr.InsertedIDs)
	}

}
func find() {
	initDb2()
	ctx := context.TODO()
	// 关闭
	defer client.Disconnect(ctx)
	c := client.Database("golang_db").Collection("Student")
	c2, err := c.Find(ctx, bson.D{{"name", "test2"}})
	if err != nil {
		log.Fatal(err)
	}
	defer c2.Close(ctx)
	for c2.Next(ctx) {
		var result bson.D
		c2.Decode(&result)
		fmt.Printf("result:%v\n", result)
		fmt.Println("*********************************")
		fmt.Printf("resultMap:%v\n", result.Map())
	}
}
func update() {
	initDb2()
	c := client.Database("golang_db").Collection("Student")
	ctx := context.TODO()
	update := bson.D{{"$set", bson.D{{"name", "big kite"}, {"age", 22}}}}
	ur, err := c.UpdateMany(ctx, bson.D{{"name", "test"}}, update)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("ur.ModifiedCount:%v\n", ur.ModifiedCount)
	}
}
func del() {
	initDb2()
	c := client.Database("golang_db").Collection("Student")
	ctx := context.TODO()
	dr, err := c.DeleteOne(ctx, bson.D{{"age", 21}})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("dr.DeletedCount:%v\n", dr.DeletedCount)
	}
}
func main() {
	// filter 过滤器 sql where
	//d := bson.D{{"name", "tom"}}
	//fmt.Printf("d:%v\n", d)

	// 插入
	//insert()
	//InsertMoney()

	// 查找
	//find()

	// 更新
	//update()

	// 删除
	del()

}
