package serializer

import (
	"github.com/hoangnguyen-1312/faucet/domain"
)

type JSONResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Error  struct {
		Code    int         `json:"code"`
		Message interface{} `json:"message"`
		Error   interface{} `json:"error"`
	} `json:"error"`
}

const (
	STATUS_FAILED  = 0
	STATUS_SUCCESS = 1
)

func NewResponseSuccess(data interface{}) JSONResponse {
	return JSONResponse{
		Status: STATUS_SUCCESS,
		Data:   data,
	}
}

func NewResponseError(err *domain.Error) JSONResponse {
	res := JSONResponse{
		Status: STATUS_FAILED,
	}
	res.Error.Code = err.Code
	res.Error.Message = err.Message
	if err.Err != nil {
		res.Error.Error = err.Err.Error()
	}
	return res
}