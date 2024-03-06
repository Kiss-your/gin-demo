package response

import "github.com/gin-gonic/gin"

func Response(c *gin.Context, httpStatus int, code string, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

func Success(c *gin.Context, data gin.H, msg string) {
	Response(c, 200, "200", data, msg)
}

func Fail(c *gin.Context, data gin.H, msg string) {
	Response(c, 400, "400", data, msg)
}
