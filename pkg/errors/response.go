package errors

import (
	json "github.com/json-iterator/go"
	"net/http"
)

type Response struct {
	Status       int    `json:"-"`
	Code         int    `json:"code,omitempty"`
	Msg          string `json:"msg,omitempty"`
	Data         any    `json:"data,omitempty"`
	EncodeMethod int    `json:"encodeMethod,omitempty"`
}

// Error
// @date 2022-09-10 17:35:32
func (r *Response) Error() string {
	if r.Data == nil {
		return r.Msg
	}

	v, _ := json.Marshal(r)
	return string(v)
}

// New
// @param msg
// @date 2022-09-10 17:35:31
func New(msg string) *Response {
	return &Response{Status: http.StatusOK, Msg: msg, Code: CodeError}
}

// NewWithCode
// @param code
// @param msg
// @date 2022-09-10 17:35:30
func NewWithCode(code int, msg string) *Response {
	return &Response{Code: code, Status: http.StatusOK, Msg: msg}
}

// WithMsg
// @param msg
// @date 2022-09-10 17:35:28
func (r *Response) WithMsg(msg string) *Response {
	r.Msg = msg
	return r
}

// WithCode
// @param code
// @date 2022-09-10 17:35:28
func (r *Response) WithCode(code int) *Response {
	r.Code = code
	return r
}

// WithData
// @param data
// @date 2022-09-10 17:35:27
func (r *Response) WithData(data any) *Response {
	r.Data = data
	return r
}

// NewWithStatus
// @param status
// @param msg
// @date 2022-09-10 17:35:26
func NewWithStatus(status int, msg string) *Response {
	return &Response{Code: CodeError, Status: status, Msg: msg}
}
