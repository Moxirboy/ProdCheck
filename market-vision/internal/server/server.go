package server

import (
	"fmt"
	configs "market-vision/internal/config"
	"market-vision/internal/service/usecase"
	"market-vision/pkg/logger"
	"market-vision/pkg/postgres"
	"market-vision/internal/controller/controller"

	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg    *configs.Config
	logger logger.Logger
}

func NewServer(
	cfg *configs.Config,
	logger logger.Logger,
) *Server {
	return &Server{
		cfg:    cfg,
		logger: logger,
	}
}

func (s Server) Run() error {
	
	pDB, err := postgres.DB(&s.cfg.Postgres)
	if err != nil {
		s.logger.Fatal(err)
	}
	g := gin.New()
	g.Use(gin.Recovery())
	uc := usecase.New(s.cfg, pDB, s.logger)
	s.logger.Info(s.cfg)
	controller.SetUp(&g.RouterGroup, s.cfg, uc, s.logger)
	return g.Run(fmt.Sprintf(":%d", s.cfg.Server.Port))

}
