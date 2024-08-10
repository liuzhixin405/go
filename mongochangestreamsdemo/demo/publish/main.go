package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// 设置 MongoDB 客户端
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("连接 MongoDB 失败:", err)
		return
	}
	defer client.Disconnect(context.TODO())

	// 获取数据库和集合
	collection := client.Database("testdb").Collection("items")

	// 插入数据
	for i := 1; i <= 5; i++ {
		item := bson.D{{"name", fmt.Sprintf("item%d", i)}, {"value", i}}
		_, err := collection.InsertOne(context.TODO(), item)
		if err != nil {
			fmt.Println("插入数据失败:", err)
			return
		}
		//fmt.Printf("插入数据: %+v\n", item)
		fmt.Printf("插入数据第 %d 条", i)
		time.Sleep(2 * time.Second) // 模拟一些延迟
	}
}
