package api

import (
	jwthandler "github.com/PentaGol/api_getway/api/token"
	"github.com/casbin/casbin/v2"
	
	"github.com/PentaGol/api_getway/api/handlers/v1"
	"github.com/PentaGol/api_getway/api/middileware"
	"github.com/PentaGol/api_getway/config"
	"github.com/PentaGol/api_getway/pkg/logger"
	"github.com/PentaGol/api_getway/services"

	_ "github.com/PentaGol/api_getway/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	CasbinEnforcer *casbin.Enforcer
}

// New ...
// @title           			Swagger for admin api
// @securityDefinitions.apikey 	ApiKeyAuth
// @in header
// @name Authorization
// @version        				1.0
// @description     			This is a admin service api.
// @Host localhost:8080
func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
		CasbinEnforcer: option.CasbinEnforcer,
	})

	jwt := jwthandler.JWTHandler{
		SiginKey: option.Conf.SiginKey,
		Log:      option.Logger,
	}

	router.Use(middileware.NewAuth(option.CasbinEnforcer, jwt, option.Conf))

	api := router.Group("/v1")
	// admins
	api.GET("/admin/:id", handlerV1.GetAdminById)
	api.POST("/admin", handlerV1.CreateAdmin)

	// login
	api.GET("/login/:email/:password", handlerV1.Login)

	// posts
	api.GET("/post/:id", handlerV1.GetPostById)
	api.POST("/post", handlerV1.CreatePost)
	api.PUT("/post/:id", handlerV1.UpdatePost)
	api.GET("/posts", handlerV1.GetAllPosts)
	api.DELETE("/post/:id", handlerV1.DeletePost)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
