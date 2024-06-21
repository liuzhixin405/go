package main

import (
	"database/sql"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type RedPacket struct {
	ID              int     `json:"id"`
	TotalAmount     float64 `json:"totalAmount"`
	TotalCount      int     `json:"totalCount"`
	RemainingAmount float64 `json:"remainingAmount"`
	RemainingCount  int     `json:"remainingCount"`
}

type Grab struct {
	ID          int     `json:"id"`
	RedPacketID int     `json:"redPacketId"`
	Amount      float64 `json:"amount"`
	Grabber     string  `json:"grabber"` //名字
	Timestamp   int64   `json:"timestamp"`
}

var db *sql.DB

func main() {
	var err error
	dsn := "root:123456Aa@tcp(127.0.0.1:3306)/redpacketdb"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	// 设置跨域中间件
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有来源访问，生产环境应谨慎使用
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200) // 预检请求返回200
			return
		}
		c.Next()
	})

	r.POST("/sendRedPacket", sendRedPacket)
	r.POST("/grabRedPacket", grabRedPacket)
	r.Run(":8080")
}

func sendRedPacket(c *gin.Context) {
	var redPacket RedPacket
	if err := c.ShouldBindJSON(&redPacket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	redPacket.RemainingAmount = redPacket.TotalAmount
	redPacket.RemainingCount = redPacket.TotalCount

	result, err := db.Exec("INSERT INTO red_packets (total_amount, total_count, remaining_amount, remaining_count) VALUES (?, ?, ?, ?)",
		redPacket.TotalAmount, redPacket.TotalCount, redPacket.RemainingAmount, redPacket.RemainingCount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"redPacketId": id})
}

func grabRedPacket(c *gin.Context) {
	var grab Grab
	if err := c.ShouldBindJSON(&grab); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx, err := db.Begin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var redPacket RedPacket
	err = tx.QueryRow("SELECT id, total_amount, total_count, remaining_amount, remaining_count FROM red_packets WHERE id = ? FOR UPDATE", grab.RedPacketID).Scan(
		&redPacket.ID, &redPacket.TotalAmount, &redPacket.TotalCount, &redPacket.RemainingAmount, &redPacket.RemainingCount)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "RedPacket not found"})
		return
	}

	if redPacket.RemainingCount <= 0 {
		tx.Rollback()
		c.JSON(http.StatusGone, gin.H{"error": "No more red packets"})
		return
	}

	rand.Seed(time.Now().UnixNano())
	amount := redPacket.RemainingAmount / float64(redPacket.RemainingCount)
	if redPacket.RemainingCount > 1 {
		amount = rand.Float64() * (redPacket.RemainingAmount / float64(redPacket.RemainingCount))
	}

	grab.Amount = amount
	grab.Timestamp = time.Now().Unix()

	_, err = tx.Exec("UPDATE red_packets SET remaining_count = remaining_count - 1, remaining_amount = remaining_amount - ? WHERE id = ?",
		grab.Amount, grab.RedPacketID)
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = tx.Exec("INSERT INTO grabs (red_packet_id, amount, grabber, timestamp) VALUES (?, ?, ?, ?)",
		grab.RedPacketID, grab.Amount, grab.Grabber, time.Unix(grab.Timestamp, 0))
	if err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = tx.Commit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"amount": grab.Amount})
}
