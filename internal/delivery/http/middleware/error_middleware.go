package middleware

import (
	"Yuk-Ujian/internal/constant"
	"Yuk-Ujian/internal/dto"
	"Yuk-Ujian/internal/pkg/apperror"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errLen := len(c.Errors)
		if errLen > 0 {
			err := c.Errors[len(c.Errors)-1].Err

			var appErr *apperror.AppError
			var vErr validator.ValidationErrors
			var utErr *json.UnmarshalTypeError
			var sErr *json.SyntaxError

			switch {
			case errors.As(err, &sErr):
				handleJsonSyntaxError(c, sErr)
				return
			case errors.As(err, &utErr):
				handleJsonUnmarshalTypeError(c, utErr)
				return
			case errors.As(err, &vErr):
				handleValidationError(c, vErr)
				return
			case errors.As(err, &appErr):
				c.AbortWithStatusJSON(appErr.Code, dto.ErrorResponse{
					Message: appErr.Message,
				})
				return

			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
					Message: constant.InternalServerErrorMessage,
				})
			}
		}
	}
}

func handleJsonSyntaxError(c *gin.Context, err *json.SyntaxError) {
	c.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
		Message: constant.JsonSyntaxError, Details: []string{err.Error()},
	})
}

func handleJsonUnmarshalTypeError(c *gin.Context, err *json.UnmarshalTypeError) {
	c.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
		Message: fmt.Sprintf(constant.InvalidJsonValueTypeError, err.Field),
	})
}

func handleValidationError(c *gin.Context, err validator.ValidationErrors) {
	ve := []dto.ValidationErrorResponse{}

	for _, fe := range err {
		ve = append(ve, dto.ValidationErrorResponse{
			Field:   fe.Field(),
			Message: tagToMsg(fe),
		})
	}

	c.AbortWithStatusJSON(http.StatusBadRequest, dto.ErrorResponse{
		Message: constant.ValidationError,
		Details: ve,
	})
}

func tagToMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "len":
		return fmt.Sprintf("%s must be %v characters", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("%s must not exceed %v characters", fe.Field(), fe.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %v", fe.Field(), fe.Param())
	case "lte":
		return fmt.Sprintf("%s must be lower than or equal to %v", fe.Field(), fe.Param())
	case "email":
		return fmt.Sprintf("%s has invalid email format", fe.Field())
	case "eq":
		return fmt.Sprintf("%s must be: %v", fe.Field(), fe.Param())
	case "min":
		return fmt.Sprintf("%s must be %v characters long", fe.Field(), fe.Param())
	case "dgte":
		return fmt.Sprintf("%s must be greater than or equal to %v", fe.Field(), fe.Param())
	default:
		return "invalid input"
	}
}
