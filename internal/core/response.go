package core

import "github.com/gin-gonic/gin"

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Success(c *gin.Context, status int, data interface{}) {
	c.JSON(status, Response{
		Data: data,
	})
}

func Error(c *gin.Context, status int, err error) {
	c.JSON(status, Response{
		Error: err.Error(),
	})
}
