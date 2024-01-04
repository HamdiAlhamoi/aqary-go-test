package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/aliml92/realworld-gin-sqlc/config"
	db "github.com/aliml92/realworld-gin-sqlc/db/sqlc"
	"github.com/aliml92/realworld-gin-sqlc/docs"
	"github.com/aliml92/realworld-gin-sqlc/logger"
)

type Server struct {
	config config.Config
	router *gin.Engine
	store  db.Store
	log    logger.Logger
}

func NewServer(config config.Config, store db.Store, log logger.Logger) *Server {
	var engine *gin.Engine
	if config.Environment == "test" {
		gin.SetMode(gin.ReleaseMode)
		fmt.Println("test environment detected")
		engine = gin.New()
	} else {
		engine = gin.Default()
	}
	server := &Server{
		config: config,
		router: engine,
		store:  store,
		log:    log,
	}
	return server
}

func (s *Server) MountHandlers() {
	api := s.router.Group("/api")
	api.POST("/users", s.RegisterUser)
	api.POST("/users/generateotp", s.GenerateOtp)
	api.POST("/users/verifyotp", s.ValidateOtp)
}

func (s *Server) MountSwaggerHandlers() {
	docs.SwaggerInfo.Version = "0.0.1"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.Title = "Conduit API"
	docs.SwaggerInfo.Description = "Conduit API Documentation"
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) Router() *gin.Engine {
	return s.router
}
