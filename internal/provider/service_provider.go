package provider

import (
	"Yuk-Ujian/internal/config"
	"Yuk-Ujian/internal/usecase"
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
