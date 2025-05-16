package middleware

import (
	"DRX_Test/internal/pkg/apperror"
	"DRX_Test/internal/pkg/logger"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type contextKey string

const requestIDKey contextKey = "request_id"

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set("request_id", requestID)

		start := time.Now()
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()

		param := map[string]interface{}{
			"request_id":  requestID,
			"client_ip":   clientIP,
			"user_agent":  userAgent,
			"status_code": c.Writer.Status(),
			"method":      c.Request.Method,
			"latency_ms":  time.Since(start).Milliseconds(),
			"path":        path,
		}

		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, requestIDKey, requestID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()

		if len(c.Errors) == 0 {
			logger.Log.WithFields(param).Info("incoming request")
		} else {
			errList := []error{}
			for _, ginErr := range c.Errors {
				err := ginErr.Err

				var appErr *apperror.AppError
				var vErr validator.ValidationErrors
				var utErr *json.UnmarshalTypeError
				var sErr *json.SyntaxError

				switch {
				case errors.As(err, &sErr), errors.As(err, &utErr), errors.As(err, &vErr):
					param["status_code"] = http.StatusBadRequest
					param["error_type"] = "validation_error"
					errList = append(errList, err)
				case errors.As(err, &appErr):
					param["status_code"] = appErr.Code
					param["error_type"] = "application_error"
					if appErr.OriginalError != nil {
						param["original_error"] = appErr.OriginalError.Error()
					}
					errList = append(errList, appErr)
				default:
					param["status_code"] = http.StatusInternalServerError
					param["error_type"] = "internal_error"
					errList = append(errList, err)
				}
			}

			if len(errList) > 0 {
				param["errors"] = errList
				logger.Log.WithFields(param).Error("request error")
			}
		}

	}
}
