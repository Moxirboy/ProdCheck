package controller

import (
	configs "market-vision/internal/config"
	"market-vision/internal/service/usecase"
	"market-vision/pkg/logger"
	v1 "market-vision/internal/controller/controller/http"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	url := ginSwagger.URL("swagger/doc.json")
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func SetUpHandlerV1(
	group *gin.RouterGroup,
	config *configs.Config,
	uc usecase.IUseCase,
	log logger.Logger,
) {
	v1.NewProductHandler(
		group,
		log,
		*config,
		uc.IProductUseCase(),
		uc.IViolationUseCase(),
	)
}
