package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pdf-server/core"
)

func NewRouter() http.Handler {
	engine := gin.Default()
	group := engine.Group("pdf-server/api/v1")
	routers := []func(group gin.IRouter){routerWK, routerChrome, routerTool}
	for _, router := range routers {
		router(group)
	}
	return engine
}

func routerWK(group gin.IRouter) {
	r := NewWKRouter()
	g := group.Group("wk")
	g.POST("convert", core.Wrap(r.Convert))
}

func routerChrome(group gin.IRouter) {
	r := NewChromeRouter()
	g := group.Group("chrome")
	g.POST("convert", core.Wrap(r.Convert))
}

func routerTool(group gin.IRouter) {
	r := NewToolRouter()
	g := group.Group("tool")
	g.GET("ping", core.Wrap(r.Ping))
}
