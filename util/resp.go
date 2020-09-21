package util

import (
	"encoding/json"
	"log"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func AjaxReturn(code int, message string, data interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (resp *Response) JsonBytes() []byte {
	r, err := json.Marshal(resp)
	if err != nil {
		log.Println(err.Error())
	}
	return r
}
