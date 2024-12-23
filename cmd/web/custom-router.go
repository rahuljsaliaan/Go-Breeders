package main

import "net/http"

type middlewareType func(http.Handler) http.Handler

type CustomRouter struct {
	middlewares []middlewareType
	routes      map[string]http.Handler
}

func NewCustomRouter() *CustomRouter {
	return &CustomRouter{
		middlewares: []middlewareType{},
		routes:      make(map[string]http.Handler),
	}
}

func (cr *CustomRouter) Use(middleware middlewareType) {
	cr.middlewares = append(cr.middlewares, middleware)
}

func (cr *CustomRouter) Handle(method, path string, handler http.HandlerFunc) {
	handlerWithMiddlewares := cr.applyMiddlewares(handler)
	cr.routes[method+path] = handlerWithMiddlewares
}

// ! This implicitly converts the HandlerFunc to a Handler by making it a struct with a ServeHTTP method
func (cr *CustomRouter) applyMiddlewares(handler http.Handler) http.Handler {
	for i := len(cr.middlewares) - 1; i >= 0; i-- {
		handler = cr.middlewares[i](handler)
	}
	return handler
}

func (cr *CustomRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, ok := cr.routes[req.Method+req.URL.Path]
	if !ok {
		http.NotFound(w, req)
		return
	}
	handler.ServeHTTP(w, req)
}

// HTTP Methods
func (cr *CustomRouter) Get(path string, handler http.HandlerFunc) {
	cr.Handle(http.MethodGet, path, handler)
}
