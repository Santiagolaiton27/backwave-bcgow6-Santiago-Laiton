package web

import "strconv"

type Response struct {
	Code    string      `json:"code"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func NewResponse(code int, data interface{}, message string) Response {
	if code < 300 {
		return Response{strconv.FormatInt(int64(code), 10), data, ""}
	}
	return Response{strconv.FormatInt(int64(code), 10), nil, message}
}
