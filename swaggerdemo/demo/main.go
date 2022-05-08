package main

import (
	"controler/controller"
	_ "controler/docs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger  demo service API
// @version 1.0
// @description This is demo server.
// @termsOfService demo.com

// @contact.name API Support
// @contact.url http://demo.com/support

// @host localhost:8091
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	r := setupRouter()
	_ = r.Run(":8091")
	log.Println("router")
}

func setupRouter() *gin.Engine {

	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Welcome To Sample program swagger"})
	})

	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/account")
		{
			accounts.POST("/create", controller.CreateAccount)
			accounts.PATCH("/update/:id", controller.UpdateAccount)
			accounts.DELETE("/delete/:id", controller.DeleteAccount)

		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r

}
