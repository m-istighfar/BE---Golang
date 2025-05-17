package validator

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

			if name == "-" {
				return ""
			}

			return name
		})

		v.RegisterCustomTypeFunc(func(field reflect.Value) interface{} {
			if valuer, ok := field.Interface().(decimal.Decimal); ok {
				return valuer.String()
			}
			return nil
		}, decimal.Decimal{})

		v.RegisterValidation("dgte", func(fl validator.FieldLevel) bool {
			data, ok := fl.Field().Interface().(string)
			if !ok {
				return false
			}
			value, err := decimal.NewFromString(data)
			if err != nil {
				return false
			}
			baseValue, err := decimal.NewFromString(fl.Param())
			if err != nil {
				return false
			}
			return value.GreaterThanOrEqual(baseValue)
		})

		v.RegisterValidation("dlte", func(fl validator.FieldLevel) bool {
			data, ok := fl.Field().Interface().(string)
			if !ok {
				return false
			}
			value, err := decimal.NewFromString(data)
			if err != nil {
				return false
			}
			baseValue, err := decimal.NewFromString(fl.Param())
			if err != nil {
				return false
			}
			return value.LessThanOrEqual(baseValue)
		})
	}
}
