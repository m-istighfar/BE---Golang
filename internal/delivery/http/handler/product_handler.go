package handler

import (
	"DRX_Test/internal/constant"
	"DRX_Test/internal/dto/pagedto"
	"DRX_Test/internal/dto/productdto"
	"DRX_Test/internal/pkg/ginutils"
	"DRX_Test/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase: productUsecase}
}

func (h *ProductHandler) GetAllProducts(ctx *gin.Context) {
	var pageInfo pagedto.PageSortDto

	if err := ctx.ShouldBindQuery(&pageInfo); err != nil {
		ctx.Error(err)
		return
	}

	if pageInfo.Page <= 0 {
		pageInfo.Page = constant.DefaultPage
	}

	if pageInfo.Limit <= 0 {
		pageInfo.Limit = constant.DefaultLimit
	}

	products, pageInfoData, err := h.productUsecase.List(ctx, &pageInfo)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutils.ResponseOKPaginated(ctx, products, *pageInfoData)
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var request productdto.CreateProductRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Error(err)
		return
	}

	product, err := h.productUsecase.Create(ctx, &request)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutils.ResponseOKData(ctx, product)
}
