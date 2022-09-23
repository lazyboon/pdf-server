package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Http struct {
	server *http.Server
}

func NewHttp(handler http.Handler) *Http {
	return &Http{
		server: &http.Server{
			Addr:    ":18080",
			Handler: handler,
		}}
}

func (h *Http) ListenAndServe() {
	go func() {
		_ = h.server.ListenAndServe()
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSEGV)
	<-c
	h.Shutdown()
}

func (h *Http) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_ = h.server.Shutdown(ctx)
	cancel()
}
