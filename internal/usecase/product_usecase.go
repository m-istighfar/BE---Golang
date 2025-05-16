package usecase

import (
	"DRX_Test/internal/dto/pagedto"
	"DRX_Test/internal/dto/productdto"
	"DRX_Test/internal/mapper"
	"DRX_Test/internal/pkg/apperror"
	"DRX_Test/internal/pkg/apputils"
	"DRX_Test/internal/repository"
	"context"
)

type ProductUsecase interface {
	List(ctx context.Context, pageInfo *pagedto.PageSortDto) ([]productdto.ProductResponse, *pagedto.PageInfoDto, error)
	Create(ctx context.Context, req *productdto.CreateProductRequest) (*productdto.ProductResponse, error)
}

type productUsecase struct {
	productRepository repository.ProductRepository
}

func NewProductUsecase(productRepository repository.ProductRepository) ProductUsecase {
	return &productUsecase{productRepository: productRepository}
}

func (uc *productUsecase) List(ctx context.Context, pageInfo *pagedto.PageSortDto) ([]productdto.ProductResponse, *pagedto.PageInfoDto, error) {
	products, totalRow, err := uc.productRepository.List(ctx, pageInfo)
	if err != nil {
		return nil, nil, apputils.HandleError(err, apperror.ErrFailedToGetProducts)
	}

	responses := mapper.ToResponses(products)

	hasNext := int64(pageInfo.Page*pageInfo.Limit) < totalRow
	hasPrev := pageInfo.Page > 1

	return responses, &pagedto.PageInfoDto{
		Page:     pageInfo.Page,
		Limit:    pageInfo.Limit,
		TotalRow: totalRow,
		HasNext:  hasNext,
		HasPrev:  hasPrev,
	}, nil
}

func (uc *productUsecase) Create(ctx context.Context, req *productdto.CreateProductRequest) (*productdto.ProductResponse, error) {

	exists, err := uc.productRepository.IsExist(ctx, req.Name)
	if err != nil {
		return nil, apputils.HandleError(err, apperror.ErrFailedToCheckProduct)
	}
	if exists {
		return nil, apperror.ErrProductAlreadyExists
	}

	product := mapper.RequestToEntity(req)
	createdProduct, err := uc.productRepository.Create(ctx, product)
	if err != nil {
		return nil, apputils.HandleError(err, apperror.ErrFailedToCreateProduct)
	}

	resp := mapper.ToResponse(createdProduct)
	return &resp, nil
}
