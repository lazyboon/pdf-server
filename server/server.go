package server

import "pdf-server/router"

func Run() {
	NewHttp(router.NewRouter()).ListenAndServe()
}

