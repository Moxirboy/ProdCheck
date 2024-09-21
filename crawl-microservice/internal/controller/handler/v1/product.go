package v1

import (
	"crawl-microservice/internal/configs"
	"crawl-microservice/internal/dto"
	"crawl-microservice/internal/models"
	"crawl-microservice/internal/service/usecase"
	"crawl-microservice/pkg/logger"
	"crawl-microservice/pkg/utils"
	"github.com/gin-gonic/gin"
	
)

type ProductHandler struct {
	log    logger.Logger
	uc     usecase.IProductUseCase
	config *configs.Config
}

// NewProductHandler initializes the product routes.
//
//	@BasePath	/api/v1
func NewProductHandler(
	r *gin.RouterGroup,
	log logger.Logger,
	config *configs.Config,
	uc usecase.IProductUseCase,
) {
	handler := &ProductHandler{
		log:    log,
		uc:     uc,
		config: config,
	}

	products := r.Group("/product")
	products.POST("", handler.create)
	products.GET("/get-all", handler.GetList)
	products.GET("/get", handler.GetByID)
}

// create creates a new product.
//
//	@Summary		Create a new product
//	@Description	Uploads an image and creates a new product record.
//	@Tags			product
//	@Accept			multipart/form-data
//	@Produce		json
//
// @Security Bearer
//
//	@Param			image		formData	file				true	"Product image"
//	@Param			name		formData	string				true	"Product name"
//	@Param			description	formData	string				false	"Product description"
//	@Param			price		formData	number				true	"Product price"
//	@Success		200			{object}	dto.BaseResponse	"Product created successfully"
//	@Failure		400			{object}	dto.BaseResponse	"Invalid request"
//	@Failure		500			{object}	dto.BaseResponse	"Internal server error"
//	@Router			/product [post]
func (h *ProductHandler) create(c *gin.Context) {
	// Handle file upload
	res:=dto.Product{}
	err := c.ShouldBind(&res)
	if err != nil {
		h.log.Error(err)
		utils.SendResponse(c, nil, err)
		return
	}
	// Validate the data
	invalidParams := utils.Validate(res)
	if invalidParams != nil {
		h.log.Error(err)
		utils.SendResponse(c, invalidParams, nil)
		return
	}


	product := models.NewProduct(res)
	if err := h.uc.Create(c.Request.Context(), product); err != nil {
		h.log.Error(err)
		utils.SendResponse(c, nil, err)
		return
	}

	utils.SendResponse(c, product, nil)
}

// GetByID retrieves a product by its ID.
//
//	@Summary		Get a product by ID
//	@Description	Retrieves a product by its unique ID.
//	@Tags			product
//	@Accept			json
//	@Produce		json
//
// @Security Bearer
//
//	@Param			id	query		string				true	"Product ID"
//	@Success		200	{object}	dto.BaseResponse	"Product retrieved successfully"
//	@Failure		400	{object}	dto.BaseResponse	"Invalid request"
//	@Failure		500	{object}	dto.BaseResponse	"Internal server error"
//	@Router			/product/get [get]
func (h *ProductHandler) GetByID(c *gin.Context) {
	id := c.Query("id")
	product, err := h.uc.GetByID(c.Request.Context(), id)
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}
	utils.SendResponse(c, ToProductResponse(*product), nil)
}

// GetList retrieves a list of products.
//
//	@Summary		Get a list of products
//	@Description	Retrieves a list of products with optional pagination and filtering by name.
//	@Tags			product
//	@Accept			json
//	@Produce		json
//
// @Security Bearer
//
//	@Param			name		query		string				false	"Product name filter"
//	@Param			page		query		int					false	"Page number"
//	@Param			page_size	query		int					false	"Page size"
//	@Success		200			{object}	dto.BaseResponse	"List of products"
//	@Failure		400			{object}	dto.BaseResponse	"Invalid request"
//	@Failure		500			{object}	dto.BaseResponse	"Internal server error"
//	@Router			/product/get-all [get]
func (h *ProductHandler) GetList(c *gin.Context) {
	name := c.Query("name")
	pageQuery, err := utils.GetPaginationFromCtx(c)
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}
	list, err := h.uc.GetList(c.Request.Context(), name, *pageQuery)
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}

	utils.SendResponse(c, list, err)
}


func (h *ProductHandler) GetListViolations(c *gin.Context) {
	name := c.Query("name")
	pageQuery, err := utils.GetPaginationFromCtx(c)
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}
	list, err := h.uc.GetListViolations(c.Request.Context(), name, *pageQuery)
	if err != nil {
		utils.SendResponse(c, nil, err)
		return
	}

	utils.SendResponse(c, list, err)
}