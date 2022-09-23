package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pdf-server/core"
)

type ToolRouter struct{}

func NewToolRouter() *ToolRouter {
	return &ToolRouter{}
}

func (t *ToolRouter) Ping(ctx *core.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": "pong",
	})
}
