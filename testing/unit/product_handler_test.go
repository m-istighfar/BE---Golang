package unit

import (
	"DRX_Test/internal/delivery/http/handler"
	"DRX_Test/internal/dto/pagedto"
	"DRX_Test/internal/dto/productdto"
	"DRX_Test/internal/pkg/validator"
	"DRX_Test/internal/usecase/mocks"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func createTestContext(method string, url string, body string) (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	if body != "" {
		ctx.Request = httptest.NewRequest(method, url, strings.NewReader(body))
		ctx.Request.Header.Set("Content-Type", "application/json")
	} else {
		ctx.Request = httptest.NewRequest(method, url, nil)
	}

	return ctx, w
}

func TestGetAllProducts(t *testing.T) {
	mockProductUsecase := mocks.NewMockProductUsecase(t)
	handler := handler.NewProductHandler(mockProductUsecase)

	mockProductUsecase.EXPECT().
		List(mock.Anything, mock.Anything).
		Return([]productdto.ProductResponse{}, &pagedto.PageInfoDto{}, nil)

	ctx, w := createTestContext("GET", "/v1/products?page=1&limit=10", "")
	handler.GetAllProducts(ctx)

	assert.Equal(t, 200, w.Code)
}

func TestCreateProduct(t *testing.T) {
	validator.RegisterValidators()
	mockProductUsecase := mocks.NewMockProductUsecase(t)
	handler := handler.NewProductHandler(mockProductUsecase)

	requestBody := productdto.CreateProductRequest{
		Name:  "Product Test",
		Price: decimal.NewFromInt(10000),
		Stock: 10,
	}

	mockProductUsecase.EXPECT().
		Create(mock.Anything, &requestBody).
		Return(&productdto.ProductResponse{}, nil)

	ctx, w := createTestContext("POST", "/v1/products", `{"name":"Product Test","price":10000,"stock":10}`)
	handler.CreateProduct(ctx)

	assert.Equal(t, 200, w.Code)
}
