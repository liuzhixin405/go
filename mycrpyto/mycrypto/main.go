// main.go
package main

import (
	"fmt"
	"log"
	"sync"
	"time"

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

	var wg sync.WaitGroup
	orders := make(chan order.Order)
	trades := make(chan trade.Trade)

	wg.Add(1)
	go generateOrders(7, orders, orderService, &wg)

	wg.Add(1)
	go processTrades(trades, tradeService, &wg)
	wg.Add(1)
	go readTradesFromConsole(trades, &wg)
	time.Sleep(10 * time.Second)
	close(orders)
	wg.Wait()
	close(trades)
	fmt.Println("Done")
}

func generateOrders(userID int, orders chan<- order.Order, orderservice order.OrderService, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-time.After(time.Second * 3):
			orderID, err := orderservice.CreateOrder(userID, "BTC/USDT", "ETH/USDT", 100.0)
			if err != nil {
				log.Println(err)
				continue
			}
			order := order.Order{ID: orderID,
				UserID:       userID,
				FromCurrency: "BTC/USDT",
				ToCurrency:   "ETH/USDT",
				Amount:       100.0,
				Status:       order.PENDING,
			}
			orders <- order
			orderID++
		default:
			return
		}
	}
}

func processTrades(trades <-chan trade.Trade, tradeService trade.TradeService, wg *sync.WaitGroup) {
	defer wg.Done()
	for trade := range trades {
		log.Println("Processing trade: ", trade)
		_, err := tradeService.ExecuteTrade(trade.UserID, trade.OrderID, trade.Amount, string(trade.Direction))
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("Trade executed successfully")
	}
}

func readTradesFromConsole(trades chan<- trade.Trade, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		var trade trade.Trade
		fmt.Println("Enter trade details: ")
		fmt.Print("user id:")
		_, err := fmt.Scanln(&trade.UserID)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Print("order id:")
		_, err = fmt.Scanln(&trade.OrderID)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Print("amount:")
		_, err = fmt.Scanln(&trade.Amount)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Print("direction (buy/sell):")
		_, err = fmt.Scanln(&trade.Direction)
		if err != nil {
			log.Println(err)
			continue
		}
		trades <- trade
	}
}
