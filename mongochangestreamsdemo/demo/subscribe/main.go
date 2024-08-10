package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	// 设置 MongoDB 客户端mongo单机模式不支持这种监听 单机报错 2024/08/10 11:18:54 (Location40573) The $changeStream stage is only supported on replica sets
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	// 获取数据库和集合
	collection := client.Database("testdb").Collection("items")

	// 设置 Change Stream
	pipeline := mongo.Pipeline{}
	changeStreamOptions := options.ChangeStream().SetFullDocument(options.UpdateLookup)
	changeStream, err := collection.Watch(context.TODO(), pipeline, changeStreamOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer changeStream.Close(context.TODO())

	fmt.Println("开始监听 Change Stream...")

	// 读取 Change Stream
	for changeStream.Next(context.TODO()) {
		var changeEvent bson.M
		if err := changeStream.Decode(&changeEvent); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("检测到更改: %+v\n", changeEvent)
	}

	if err := changeStream.Err(); err != nil {
		log.Fatal(err)
	}
}
