package handler

import (
	_ "crawl-microservice/docs"
	"crawl-microservice/internal/configs"
	v1 "crawl-microservice/internal/controller/handler/v1"
	"crawl-microservice/internal/service/storage/repo"
	"crawl-microservice/internal/service/usecase"
	"crawl-microservice/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUp(
	g *gin.RouterGroup,
	config *configs.Config,
	uc usecase.IUseCase,
	log logger.Logger,
	redis repo.IRedisRepository,
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
	v1.NewAuthHandler(
		group,
		log,
		uc.IAuthUseCase(),
		uc.IAccountUseCase(),
	)
	v1.NewProductHandler(
		group,
		log,
		config,
		uc.ProductUsecase(),
	)
	
}
