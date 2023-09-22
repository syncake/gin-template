package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ApiCode int

var (
	ApiSuccessful ApiCode = 0
	ApiFailure    ApiCode = 500
)

type Response struct {
	Code   ApiCode                 `json:"code"`
	Msg    string                  `json:"msg"`
	Result *map[string]interface{} `json:"result"`
}

func (resp *Response) String() string {
	_json, err := json.Marshal(resp)
	if err != nil {
		fmt.Print(err)
	}

	return string(_json)
}

func Success(c *gin.Context, msg string, data *map[string]interface{}) {
	if msg == "" {
		msg = "success"
	}

	JSON(c, ApiSuccessful, msg, data)
	return
}

func Fail(c *gin.Context, msg string, code ApiCode, data *map[string]interface{}) {
	if code == 0 {
		code = ApiFailure
	}

	if msg == "" {
		msg = "failed"
	}

	JSON(c, code, msg, data)
	return
}

func JSON(
	c *gin.Context,
	code ApiCode,
	msg string,
	data *map[string]interface{},
) {
	c.JSON(http.StatusOK, &Response{
		Code:   code,
		Msg:    msg,
		Result: data,
	})
	return
}
