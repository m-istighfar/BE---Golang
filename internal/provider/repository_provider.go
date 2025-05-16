package provider

import (
	"Yuk-Ujian/internal/pkg/database"
	"Yuk-Ujian/internal/repository"
)

type Repositories struct {
	Product repository.ProductRepository
}

func ProvideRepositories(pg *database.PostgresWrapper) *Repositories {
	return &Repositories{
		Product: repository.NewProductRepository(pg),
	}
}
