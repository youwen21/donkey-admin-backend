package appresp

import (
	"errors"
	"gofly/apperror"
)

type RespData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Ok() *RespData {
	return &RespData{
		Code: 0,
		Msg:  "ok",
	}
}

func Data(data interface{}) *RespData {
	return &RespData{
		Code: 0,
		Msg:  "ok",
		Data: data,
	}
}

func CodeErr(code int, err error) *RespData {
	return &RespData{
		Code: code,
		Msg:  err.Error(),
	}
}

func extractErr(err error) (int, string) {
	if err == nil {
		return 0, "ok"
	}

	code := 9999

	var appErr apperror.AppError
	ok := errors.As(err, &appErr)
	if ok {
		code = appErr.Code
	}
	return code, err.Error()
}

func Err(err error) *RespData {
	if err == nil {
		return Ok()
	}

	code, msg := extractErr(err)

	return &RespData{
		Code: code,
		Msg:  msg,
	}
}

func ErrMsg(errMsg string) *RespData {
	return &RespData{
		Code: 9999,
		Msg:  errMsg,
	}
}

func ErrData(data interface{}, err error) *RespData {
	if err != nil {
		return Data(data)
	}

	code, msg := extractErr(err)

	return &RespData{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func Reps(data interface{}, err error) *RespData {
	if err != nil {
		return Err(err)
	}
	return Data(data)
}

func Raw(code int, data interface{}, err error) *RespData {
	return &RespData{
		Code: code,
		Msg:  err.Error(),
		Data: data,
	}
}
