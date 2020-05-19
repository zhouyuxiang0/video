package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func newMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check session
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func registerHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", createUser)
	router.POST("/user/:user_name", login)
	return router
}

func main() {
	r := registerHandlers()
	middleWareHandler := newMiddleWareHandler(r)
	http.ListenAndServe(":8000", middleWareHandler)
}
