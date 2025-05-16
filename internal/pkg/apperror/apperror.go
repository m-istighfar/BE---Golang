package apperror

import (
	"DRX_Test/internal/constant"
	"net/http"
)

type AppError struct {
	Code          int    `json:"code"`
	Message       string `json:"message"`
	OriginalError error  `json:"-"`
}

func (ce *AppError) Error() string {
	return ce.Message
}

var (
	ErrInternalServerError   = &AppError{http.StatusInternalServerError, constant.InternalServerErrorMessage, nil}
	ErrFailedToGetProducts   = &AppError{http.StatusInternalServerError, "failed to get products", nil}
	ErrFailedToCreateProduct = &AppError{http.StatusInternalServerError, "failed to create product", nil}
	ErrFailedToCheckProduct  = &AppError{http.StatusInternalServerError, "failed to check product", nil}
	ErrProductAlreadyExists  = &AppError{http.StatusBadRequest, "product already exists", nil}
)
