package api

import (
	"fmt"
	"infra-base-go/internal/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	server *echo.Echo
	cfg    *config.Config
}

func New(db *gorm.DB, cfg *config.Config) *Server {

	// add middlewares
	server := echo.New()

	server.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			fmt.Printf("REQUEST: uri: %v, status: %v, custom-value: %v\n", v.URI, v.Status)
			return nil
		},
	}))
	server.Use(middleware.Recover())
	server.Use(middleware.CORS())

	return &Server{
		server: server,
		db:     db,
		cfg:    cfg,
	}
}

func (s *Server) SetupRoutes() {
	api := s.server.Group("api/v1")
	{
		authGroup := api.Group("/auth")
		{
			authGroup.POST("/register", func(c echo.Context) error {
				return nil
			})
			authGroup.POST("/login", func(c echo.Context) error {
				return nil
			})
		}

	}
}

func (s *Server) Start() error {
	return s.server.Start(":" + s.cfg.Server.Port)
}

func (s *Server) ShutDown() error {
	return s.server.Close()
}
