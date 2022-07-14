package app

import (
	"github.com/QinLiStudio/Conship/internal/app/config"
	"github.com/QinLiStudio/Conship/internal/app/middleware"
	"github.com/QinLiStudio/Conship/internal/app/router"
	"github.com/gin-gonic/gin"
)

func InitGinEngine(r router.IRouter) *gin.Engine {
	gin.SetMode(config.C.RunMode)

	app := gin.New()

	// prefixes := r.Prefixes()

	// 跨域处理
	if config.C.CORS.Enable {
		app.Use(middleware.CORSMiddleware())
	}

	r.Register(app)

	return app
}
