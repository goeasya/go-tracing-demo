package httpx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SuccessCode  = 0
	EventErrCode = 100001
	MsgSuccess   = "success"
)

//ResponseBody
type ResponseBody struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Msg     string      `json:"msg"`
}

// GinResponseSuccess
func GinResponseSuccess(c *gin.Context, code int, data interface{}) {
	if data == nil {
		data = struct{}{}
	}
	result := ResponseBody{
		Code:    code,
		Success: true,
		Data:    data,
		Msg:     MsgSuccess,
	}
	c.SecureJSON(http.StatusOK, result)
}

// GinResponseError
func GinResponseError(c *gin.Context, code int, msg string) {
	result := ResponseBody{
		Code:    code,
		Success: false,
		Data:    struct{}{},
		Msg:     msg,
	}
	c.SecureJSON(http.StatusInternalServerError, result)
}