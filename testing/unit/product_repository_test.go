package unit

import (
	"DRX_Test/internal/dto/pagedto"
	"DRX_Test/internal/pkg/database"
	"DRX_Test/internal/repository"
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestProductRepository_List(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock: %v", err)
	}
	defer db.Close()

	pgWrapper := database.NewPostgresWrapper(db)

	now := time.Now()

	pageInfo := &pagedto.PageSortDto{
		Page:   1,
		Limit:  10,
		Search: "test",
		SortBy: "name",
	}

	mock.ExpectQuery(`SELECT.*`).
		WithArgs("%test%", 10, 0).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "price", "stock", "created_at", "updated_at", "total_row"}).
				AddRow(1, "Product 1", 10000, 10, now, now, 1),
		)

	repo := repository.NewProductRepository(pgWrapper)
	products, total, err := repo.List(context.Background(), pageInfo)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), total)
	assert.Len(t, products, 1)

	p := products[0]
	assert.Equal(t, 1, p.ID)
	assert.Equal(t, "Product 1", p.Name)
	assert.True(t, decimal.NewFromInt(10000).Equal(p.Price))
	assert.Equal(t, 10, p.Stock)
	assert.True(t, now.Equal(p.CreatedAt))
	assert.True(t, now.Equal(p.UpdatedAt))

	assert.NoError(t, mock.ExpectationsWereMet())
}
