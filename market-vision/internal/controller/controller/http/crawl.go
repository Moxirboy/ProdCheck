package v1

import (
	"context"
	"log"
	configs "market-vision/internal/config"
	"market-vision/internal/controller/dto"
	"market-vision/internal/service/usecase"
	"market-vision/pkg/logger"
	rabbitmq "market-vision/pkg/rabbitMq"
	"market-vision/prod-business/crawl"

	"github.com/gin-gonic/gin"
)

// ProductHandler handles authentication-related operations.
type ProductHandler struct {
	log logger.Logger
	cfg configs.Config
	uc  usecase.IProductUseCase
	ucv usecase.IViolationUseCase
}

// NewProductHandler initializes the authentication routes.
//
//	@BasePath	/api/v1
func NewProductHandler(
	r *gin.RouterGroup,
	l logger.Logger,
	cfg configs.Config,
	uc  usecase.IProductUseCase,
	ucv usecase.IViolationUseCase,

) {
	handler := &ProductHandler{
		log: l,
		cfg: cfg,
		uc:  uc,
		ucv: ucv,
	}
	products := r.Group("/product")
	products.POST("/details", handler.GetProductDetails)
}


func (h *ProductHandler) GetProductDetails(c *gin.Context) {
	res := dto.ProductIncome{}
	err := c.ShouldBindJSON(&res)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"message": "Invalid request"})
		return
	}
	product, err := crawl.Crawl(h.cfg, *dto.FromProductIncome(res))
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"message": "Invalid request"})
		return
	}
	details,err:=h.uc.GetProductDetails(context.Background(), res.ProductID)
	if err!=nil{
		log.Println(err)
		c.JSON(400, gin.H{"message": "Invalid request"})
		return
	}
	id,err:=h.uc.CreateProduct(context.Background(),product)
	if details.Map<product.GetAdvertaisedPrice(){
		Id,err:=h.ucv.CreateViolation(context.Background(),id)
		if err!=nil{
			log.Println(err)
			c.JSON(400, gin.H{"message": "Invalid request"})
			return
		}
		rabbitmq.Publish(Id+"/"+res.Url,"violation")
	}
	c.JSON(200, gin.H{"message": "Success"})
}
