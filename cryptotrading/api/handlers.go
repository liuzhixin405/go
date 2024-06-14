package api

import (
	"context"
	"cryptotrading/pb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/trade", func(c *gin.Context) {
		var req pb.TradeRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		client := pb.NewTradeServiceClient(conn)
		grpcReq, err := client.ExecuteTrade(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, grpcReq)
	})

	return r
}
