package provider

import (
	"DRX_Test/internal/config"
	"DRX_Test/internal/usecase"
)

type Usecases struct {
	Product usecase.ProductUsecase
}

func ProvideUsecases(configs *config.Config, repositories *Repositories) *Usecases {

	ProductUsecase := usecase.NewProductUsecase(repositories.Product)

	return &Usecases{
		Product: ProductUsecase,
	}
}
