package routers

import (
	"rices/apis/controllers"
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
	user *controllers.UserController,
	auth *middlewares.MiddlewareJwt,

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

	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", user.Register)
		userGroup.POST("/login", user.Login)
		authorized := userGroup.Group("/")
		authorized.Use(auth.Authorization())
		{
			authorized.GET("/profile", user.Profile)
		}
	}
	return &ApiRouter{
		Engine: engine,
	}
}
