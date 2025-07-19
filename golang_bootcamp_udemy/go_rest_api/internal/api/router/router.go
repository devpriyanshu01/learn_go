package router

import (
	"net/http"
)

func MainRouter() *http.ServeMux {
	tRouter := teacherRouter()
	sRouter := studentRouter()
	eRouter := execsRouter()

	tRouter.Handle("/", sRouter)
	sRouter.Handle("/", eRouter)
	return tRouter
}
