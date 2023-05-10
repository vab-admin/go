package errors

import (
	"net/http"
)

const (
	CodeOK    int = 200
	CodeError int = 100000
)

var (
	ErrInvalidParam   = New("参数有误")
	ErrAPINotFound    = NewWithStatus(http.StatusNotFound, "不存在的接口")
	ErrInternalServer = NewWithCode(http.StatusInternalServerError, "服务器繁忙")
	ErrCanceled       = New("取消操作")

	ErrAdminUserAccountNotFound = New("不存在的账号")
	ErrAdminUserPasswordInvalid = New("密码有误")
)
