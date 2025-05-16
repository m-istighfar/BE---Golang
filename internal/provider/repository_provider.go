package provider

import (
	"DRX_Test/internal/pkg/database"
	"DRX_Test/internal/repository"
)

type Repositories struct {
	Product repository.ProductRepository
}

func ProvideRepositories(pg *database.PostgresWrapper) *Repositories {
	return &Repositories{
		Product: repository.NewProductRepository(pg),
	}
}
