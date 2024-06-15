// main.go
package main

import (
	"fmt"
	"log"

	"mycrypto/order"
	"mycrypto/trade"
	"mycrypto/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 连接 MySQL 数据库
	dsn := "root:123456@tcp(127.0.0.1:3307)/crypto?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// 自动迁移以创建表
	err = db.AutoMigrate(&user.User{}, &trade.Trade{}, &order.Order{})
	if err != nil {
		log.Fatal("failed to migrate database schema: ", err)
	}

	log.Println("Database migrated successfully")

	// 创建 MySQL 实现的用户、订单和交易存储库
	userRepo := user.NewUserRepository(db)
	orderRepo := order.NewMySqlUserRepository(db)

	tradeRepo := trade.NewMySqlUserRepository(db)
	orderService := order.NewOrderService(orderRepo)
	// 创建用户、订单和交易服务
	userService := user.NewUserService(userRepo)

	tradeService := trade.NewTradeService(tradeRepo, orderRepo, userService)
	login, err := userService.Login("user1", "password")
	if !login || err != nil {
		log.Fatal("Failed to login: ", err)
	}

	// 模拟用户注册和下单
	userID, err := userService.Register("user6", "user1@example6.com", "password")
	if err != nil {
		log.Fatalf("Failed to register user: %v", err)
	}

	orderID, err := orderService.CreateOrder(userID, "USD", "EUR", 100.0)
	if err != nil {
		log.Fatalf("Failed to create order: %v", err)
	}

	// 模拟执行交易
	tradeID, err := tradeService.ExecuteTrade(userID, orderID, 100.0, "Buy")
	if err != nil {
		log.Fatalf("Failed to execute trade: %v", err)
	}

	fmt.Printf("Trade executed successfully. Trade ID: %d\n", tradeID)
}
