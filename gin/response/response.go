package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpBody struct {
	Status int         `json:"status" comment:"1成功 0失败"`
	Msg    string      `json:"msg" comment:"返回信息"`
	Data   interface{} `json:"data" comment:"返回数据"`
}

const (
	SUCCESS = 1
	ERROR   = 0
)

func HttpResult(status int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, HttpBody{
		status,
		msg,
		data,
	})
}
