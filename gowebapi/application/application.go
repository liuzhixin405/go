package application

import "github.com/gin-gonic/gin"

type Application struct {
	Router *gin.Engine
}

func (app *Application) Run(addr string) {
	app.Router.Run(addr)
}
