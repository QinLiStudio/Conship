package router

import (
	"github.com/QinLiStudio/Conship/internal/app/api"
	"github.com/QinLiStudio/Conship/internal/app/configs"
	"github.com/QinLiStudio/Conship/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func Register(app *gin.Engine) error {
	RegisterAPI(app)
	return nil
}

func RegisterAPI(app *gin.Engine) {
	app.Use(middleware.LimitRoute(configs.CONFIG.Limit.Limit))
	{
		app.POST("/meta", api.Upload)
		app.PUT("/meta", api.Update)
		app.DELETE("/meta", api.Delete)
		app.GET("/meta/*path", api.Search)
	}
}
