package requesto

import (
	"MyBalance/internal/http/requesto/requesto_errors"
	"net/http"
)

const ErrCodeInternalServiceFailure = "INTERNAL_SERVICE_FAILURE"
const ErrCodeUnmarshalError = "UNMARSHAL_ERROR"
const ErrCodeTooManyRequests = "TOO_MANY_REQUESTS"
const ErrCodeBadRequest = "BAD_REQUEST"

var (
	FailedRequestCreation = requesto_errors.Error{HttpCode: http.StatusInternalServerError, ErrText: "failed to create request", ErrCode: ErrCodeInternalServiceFailure}
	ErrorUnmarshal        = &requesto_errors.Error{HttpCode: http.StatusInternalServerError, ErrText: "Unmarshal error", ErrCode: ErrCodeUnmarshalError}
	TooManyRequests       = &requesto_errors.Error{HttpCode: http.StatusInternalServerError, ErrText: "Too Many Requests", ErrCode: ErrCodeTooManyRequests}
	InternalError         = &requesto_errors.Error{HttpCode: http.StatusInternalServerError, ErrText: "Internal Service Error", ErrCode: ErrCodeInternalServiceFailure}
	UserNotFound          = &requesto_errors.Error{HttpCode: http.StatusBadRequest, ErrText: "User Not Found", ErrCode: ErrCodeBadRequest}
)
