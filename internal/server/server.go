package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/priyanshu334/gopr1/internal/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Server struct {
	App    *fiber.App
	Config *config.Config
	Logger *zap.Logger
	DB     *gorm.DB
}

func New(cfg *config.Config, log *zap.Logger, db *gorm.DB) *Server {
	App := fiber.New()
	return &Server{
		App:    App,
		Config: cfg,
		Logger: log,
		DB:     db,
	}
}

func (s *Server) Start() error {
	return s.App.Listen(":" + s.Config.AppPort)
}
