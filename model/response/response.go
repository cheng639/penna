package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Pagination struct {
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Total int64 `json:"total"`
}

type PagedData struct {
	List       interface{} `json:"list"`
	Pagination *Pagination `json:"pagination"`
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SUCCESS = 0
	ERROR   = 1
)

func Result(code int, msg string, data interface{}, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

func Success(c *gin.Context) {
	Result(SUCCESS, "success", map[string]interface{}{}, c)
}

func Message(message string, c *gin.Context) {
	Result(SUCCESS, message, map[string]interface{}{}, c)
}

func Data(data interface{}, c *gin.Context) {
	Result(SUCCESS, "OK", data, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, "failed", map[string]interface{}{}, c)
}

func Error(message string, c *gin.Context) {
	Result(ERROR, message, map[string]interface{}{}, c)
}
