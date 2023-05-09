package api

import (
	jwthandler "github.com/PentaGol/api/token"
	"github.com/PentaGol/storage"
	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"

	v1 "github.com/PentaGol/api/handlers/v1"
	"github.com/PentaGol/api/middileware"
	"github.com/PentaGol/config"
	"github.com/PentaGol/pkg/logger"

	_ "github.com/PentaGol/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// Option ...
type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	Strg           storage.IStorage
	CasbinEnforcer *casbin.Enforcer
	Jwthandler jwthandler.JWTHandler 
}

// New ...
// @title           			created by ravshan
// @securityDefinitions.apikey 	ApiKeyAuth
// @in header
// @name Authorization
// @version        				1.0
// @description     			This is a admin api.
// @Host localhost:8080
func New(option Option) *gin.Engine {
	router := gin.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "*")
	router.Use(cors.New(corsConfig))

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		Cfg:            option.Conf,
		CasbinEnforcer: option.CasbinEnforcer,
		Strg:           option.Strg,
		JwtHandler: option.Jwthandler,
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
	api.POST("/login", handlerV1.Login)

	// posts
	api.GET("/post/:id", handlerV1.GetPostById)
	api.POST("/post", handlerV1.CreatePost)
	api.PUT("/post/:id", handlerV1.UpdatePost)
	api.GET("/posts", handlerV1.GetAllPosts)
	api.DELETE("/post/:id", handlerV1.DeletePost)

	// Liga
	api.POST("/liga", handlerV1.CreateLiga)
	api.GET("/liga/:id", handlerV1.GetLigaById)
	api.GET("/ligas", handlerV1.GetAllLigas)
	api.DELETE("/liga/:id", handlerV1.DeleteLiga)

	//Game
	api.POST("/game", handlerV1.CreateGame)
	api.GET("/game/:id", handlerV1.GetGameById)
	api.GET("/games", handlerV1.GetAllGames)
	api.DELETE("/game/:id", handlerV1.DeleteGame)

	//Game
	api.POST("/club", handlerV1.CreateClub)
	api.GET("/club/:id", handlerV1.GetClubById)
	api.GET("/clubs", handlerV1.GetAllClubs)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
