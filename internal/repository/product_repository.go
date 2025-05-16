package repository

import (
	"Yuk-Ujian/internal/dto/pagedto"
	"Yuk-Ujian/internal/entity"
	"Yuk-Ujian/internal/pkg/apputils"
	"Yuk-Ujian/internal/pkg/database"
	"context"
	"fmt"
	"strings"
)

type ProductRepository interface {
	List(ctx context.Context, pageInfo *pagedto.PageSortDto) ([]*entity.Product, int64, error)
	Create(ctx context.Context, product *entity.Product) (*entity.Product, error)
	IsExist(ctx context.Context, name string) (bool, error)
}

type productRepository struct {
	db *database.PostgresWrapper
}

func NewProductRepository(db *database.PostgresWrapper) *productRepository {
	return &productRepository{db: db}
}

func (r *productRepository) List(ctx context.Context, pageInfo *pagedto.PageSortDto) ([]*entity.Product, int64, error) {
	offset := (pageInfo.Page - 1) * pageInfo.Limit
	orderBy := apputils.ConvertSortByToSQL(pageInfo.SortBy, "created_at", "DESC")

	var queryBuilder strings.Builder
	queryBuilder.WriteString(`
		SELECT
			p.id,
			p.name,
			p.price,
			p.stock,
			p.created_at,
			p.updated_at,
			COUNT(*) OVER() AS total_row
		FROM products p
		WHERE p.name ILIKE $1
	`)

	params := []interface{}{"%" + pageInfo.Search + "%"}

	queryBuilder.WriteString(fmt.Sprintf(`
		ORDER BY %s
		LIMIT $%d OFFSET $%d
	`, orderBy, len(params)+1, len(params)+2))

	params = append(params, pageInfo.Limit, offset)

	rows, err := r.db.Start(ctx).QueryContext(ctx, queryBuilder.String(), params...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var products []*entity.Product
	var totalRow int64
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.CreatedAt,
			&product.UpdatedAt,
			&totalRow,
		); err != nil {
			return nil, 0, err
		}
		products = append(products, &product)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return products, totalRow, nil
}

func (r *productRepository) Create(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	query := `
        INSERT INTO products (name, price, stock)
        VALUES ($1, $2, $3)
        RETURNING id, name, price, stock, created_at, updated_at
    `

	var createdProduct entity.Product
	err := r.db.Start(ctx).QueryRowContext(
		ctx,
		query,
		product.Name,
		product.Price,
		product.Stock,
	).Scan(
		&createdProduct.ID,
		&createdProduct.Name,
		&createdProduct.Price,
		&createdProduct.Stock,
		&createdProduct.CreatedAt,
		&createdProduct.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &createdProduct, nil
}

func (r *productRepository) IsExist(ctx context.Context, name string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM products WHERE name = $1)`

	var exists bool
	err := r.db.Start(ctx).QueryRowContext(ctx, query, name).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
