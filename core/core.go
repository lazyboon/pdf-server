package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Context struct {
	*gin.Context
}

func (c *Context) AbortWithData(data any) {
	c.Context.JSON(http.StatusOK, data)
	c.Context.Abort()
}

func (c *Context) AbortBadRequest() {
	c.AbortWithData(gin.H{
		"code": http.StatusBadRequest,
		"msg":  "bad request",
	})
}

func (c *Context) AbortWithError(err error) {
	c.AbortWithData(gin.H{
		"code":    http.StatusInternalServerError,
		"message": err.Error(),
	})
}

func Wrap(handler func(c *Context)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c := &Context{Context: ctx}
		handler(c)
		ctx.Next()
	}
}
