package mapper

import (
	"DRX_Test/internal/dto/productdto"
	"DRX_Test/internal/entity"
)

func ToResponse(p *entity.Product) productdto.ProductResponse {
	return productdto.ProductResponse{
		ID:        p.ID,
		Name:      p.Name,
		Price:     p.Price,
		Stock:     p.Stock,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func ToResponses(products []*entity.Product) []productdto.ProductResponse {
	responses := make([]productdto.ProductResponse, len(products))
	for i, p := range products {
		responses[i] = ToResponse(p)
	}
	return responses
}

func RequestToEntity(req *productdto.CreateProductRequest) *entity.Product {
	return &entity.Product{
		Name:  req.Name,
		Price: req.Price,
		Stock: req.Stock,
	}
}
