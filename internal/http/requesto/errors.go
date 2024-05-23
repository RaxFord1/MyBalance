package requesto

import (
	"MyBalance/internal/http/requesto/requesto_errors"
	"net/http"
)

const ErrCodeInternalServiceFailure = "INTERNAL_SERVICE_FAILURE"

var (
	FailedRequestCreation = requesto_errors.Error{HttpCode: http.StatusInternalServerError, ErrText: "failed to create request", ErrCode: ErrCodeInternalServiceFailure}
	ErrorUnmarshal        = &requesto_errors.Error{HttpCode: http.StatusInternalServerError, ErrCode: "unmarshal ERROR"}
)
