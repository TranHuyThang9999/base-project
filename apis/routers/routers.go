package routers

import (
	"demo_time_sheet_server/apis/controllers"
	"demo_time_sheet_server/apis/middlewares"
	"demo_time_sheet_server/common/configs"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
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
	engine.Use(cors.CorsAPI())

	p := ginprometheus.NewPrometheus("gin")
	p.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
		return c.FullPath()
	}
	p.Use(engine)
	r := engine.RouterGroup.Group("/manager")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", user.Register)
		userGroup.POST("/login", user.Login)
		userGroup.POST("/register/auth2", user.LoginWithGG)
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
