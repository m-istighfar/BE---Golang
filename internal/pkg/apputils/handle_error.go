package apputils

import "Yuk-Ujian/internal/pkg/apperror"

func HandleError(err error, appError *apperror.AppError) error {
	if err == nil {
		return nil
	}

	return &apperror.AppError{
		Code:          appError.Code,
		Message:       appError.Message,
		OriginalError: err,
	}
}
