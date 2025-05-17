package unit

import (
	"context"
	"errors"
	"testing"

	"DRX_Test/internal/dto/pagedto"
	"DRX_Test/internal/entity"
	"DRX_Test/internal/repository/mocks"
	"DRX_Test/internal/usecase"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestProductUsecase_List(t *testing.T) {
	ctx := context.Background()

	mockRepo := new(mocks.MockProductRepository)

	uc := usecase.NewProductUsecase(mockRepo)

	pageInfo := &pagedto.PageSortDto{
		Page:   1,
		Limit:  2,
		SortBy: "created_at:desc",
		Search: "",
	}

	t.Run("success", func(t *testing.T) {
		mockProducts := []*entity.Product{
			{
				ID:    1,
				Name:  "Product A",
				Price: decimal.NewFromInt(10000),
				Stock: 10,
			},
			{
				ID:    2,
				Name:  "Product B",
				Price: decimal.NewFromInt(10000),
				Stock: 5,
			},
		}
		totalRow := int64(5)

		mockRepo.On("List", ctx, pageInfo).Return(mockProducts, totalRow, nil).Once()

		result, page, err := uc.List(ctx, pageInfo)

		assert.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, int64(5), page.TotalRow)
		assert.Equal(t, true, page.HasNext)
		assert.Equal(t, false, page.HasPrev)

		mockRepo.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("List", ctx, pageInfo).Return(nil, int64(0), errors.New("db error")).Once()

		result, page, err := uc.List(ctx, pageInfo)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Nil(t, page)

		mockRepo.AssertExpectations(t)
	})
}
