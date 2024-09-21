package controller

import (
	configs "engine/internal/configs"
	v1 "engine/internal/controller/http"
	rabbitmq "engine/internal/controller/rabbitMq"
	"engine/internal/service/usecase"
	"engine/pkg/logger"

	"github.com/gin-gonic/gin"
)

func SetUp(
	g *gin.RouterGroup,
	config *configs.Config,
	uc usecase.IUseCase,
	log logger.Logger,
) {
	SetUpHandlerV1(
		g.Group("/api/v1"),
		config, uc, log,
	)
}

func SetUpHandlerV1(
	group *gin.RouterGroup,
	config *configs.Config,
	uc usecase.IUseCase,
	log logger.Logger,
) {
	v1.NewScreenshotController(
		group,
	)
	rabbitmq.NewRabbitMq(*config,uc.ScreenshotUsecase())
}
