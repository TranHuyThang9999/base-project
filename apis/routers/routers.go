package routers

import (
	"rices/apis/middlewares"
	"rices/common/configs"

	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	Engine *gin.Engine
}

func NewApiRouter(
	cf *configs.Configs,
	cors *middlewares.MiddlewareCors,
) *ApiRouter {
	engine := gin.New()
	gin.DisableConsoleColor()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(cors.Cors())
	r := engine.RouterGroup.Group("/manager")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	return &ApiRouter{
		Engine: engine,
	}
}
