package server

import (
	"fmt"
	configs "engine/internal/configs"
	"engine/internal/service/usecase"
	"engine/pkg/logger"
	"engine/pkg/postgres"
	"engine/internal/controller"

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
