package errorresponse

import (
	"encoding/json"

	errorcodes "wowcollector.io/internal/common/error-codes"
)

type ErrorResponse struct {
	Code    errorcodes.ErrorCode `json:"code"`
	Message string               `json:"message"`
}

func GenerateErrorBody(code errorcodes.ErrorCode, message string) []byte {
	body, _ := json.Marshal(ErrorResponse{
		Code:    code,
		Message: message,
	})
	return body
}
