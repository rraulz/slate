package errors

import "errors"

type APIError struct {
	Status          int
	ResponseMessage string
	InternalMessage string
}

func ConvertErrorToGinResponse(err error) APIError {
	var apiError APIError
	var svcError *ServiceError

	if errors.As(err, &svcError) {
		apiError.ResponseMessage = svcError.GetSvcError().Error()
		apiError.InternalMessage = svcError.GetOgError().Error()
		svcErr := svcError.GetSvcError()
		switch svcErr {
		case ErrBadRequest:
			apiError.Status = 400
		case ErrUnauthorized:
			apiError.Status = 401
		case ErrInternalError:
			apiError.Status = 500
		}
	}

	return apiError
}
