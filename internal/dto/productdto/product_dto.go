package productdto

import (
	"time"

	"github.com/shopspring/decimal"
)

type ProductResponse struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	Price     decimal.Decimal `json:"price"`
	Stock     int             `json:"stock"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type CreateProductRequest struct {
	Name  string          `json:"name" binding:"required"`
	Price decimal.Decimal `json:"price" binding:"required,dgte=1"`
	Stock int             `json:"stock" binding:"required,min=0"`
}
