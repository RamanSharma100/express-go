package http

import (
	"strings"
)

type Router struct {
	routes      []Route
	middlewares []Middleware
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) ValidateRoute(path string, handler Handler) bool {
	if path == "" {
		panic("Path cannot be empty")
	}
	if handler == nil {
		panic("Handler cannot be empty")
	}

	return true
}

func (r *Router) getParameterizedRoute(path string) (string, []string) {
	Params := []string{}
	parts := strings.Split(path, "/")
	for i, part := range parts {
		if strings.HasPrefix(part, ":") {
			paramName := strings.TrimPrefix(part, ":")
			Params = append(Params, paramName)
			parts[i] = "{" + paramName + "}"
		}
	}

	return strings.Join(parts, "/"), Params
}

func (r *Router) AddRoute(path string, handler Handler, method []string) {
	if r.ValidateRoute(path, handler) {

		if len(method) == 0 {
			method = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"}
		}

		params := []string{}

		if isParameterizedRoute(path) {
			path, params = r.getParameterizedRoute(path)
		}

		r.routes = append(r.routes, Route{
			Method:      method,
			Path:        path,
			Handler:     handler,
			Params:      params,
			Middlewares: append([]Middleware{}, r.middlewares...),
		})
	}
}

func (r *Router) Get(path string, handler Handler) {
	r.AddRoute(path, handler, []string{"GET"})
}

func (r *Router) Post(path string, handler Handler) {
	r.AddRoute(path, handler, []string{"POST"})
}

func (r *Router) Put(path string, handler Handler) {
	r.AddRoute(path, handler, []string{"PUT"})
}

func (r *Router) Delete(path string, handler Handler) {
	r.AddRoute(path, handler, []string{"DELETE"})
}

func (r *Router) Patch(path string, handler Handler) {
	r.AddRoute(path, handler, []string{"PATCH"})
}

func (r *Router) Options(path string, handler Handler) {
	r.AddRoute(path, handler, []string{"OPTIONS"})
}

func (r *Router) Head(path string, handler Handler) {
	r.AddRoute(path, handler, []string{"HEAD"})
}

func (r *Router) Add(path string, handler Handler) {
	r.AddRoute(path, handler, []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"})
}

func (r *Router) UseRouter(path string, router *Router) {
	if router == nil {
		return
	}

	if path == "" || path[0] != '/' {
		path = "/" + path
	}

	for _, route := range router.routes {
		if route.Path != "/" && route.Path[0] != '/' {
			route.Path = "/" + route.Path
		}
		fullPath := path + route.Path
		r.addRouteWithMiddleware(fullPath, route.Handler, route.Method, route.Middlewares...)
	}
}

func (r *Router) Use(middlewares ...Middleware) {
	if middlewares == nil {
		panic("Middleware cannot be nil")
	}
	r.middlewares = append(r.middlewares, middlewares...)
}

func (r *Router) Group(prefix string, middlewares []Middleware, handler func(router *Router)) {
	newRouter := &Router{
		routes:      []Route{},
		middlewares: middlewares,
	}

	handler(newRouter)

	r.UseRouter(prefix, newRouter)
}

func (r *Router) addRouteWithMiddleware(path string, handler Handler, method []string, middlewares ...Middleware) {
	if r.ValidateRoute(path, handler) {

		if len(method) == 0 {
			method = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"}
		}

		params := []string{}

		if isParameterizedRoute(path) {
			path, params = r.getParameterizedRoute(path)
		}

		r.routes = append(r.routes, Route{
			Method:      method,
			Path:        path,
			Handler:     handler,
			Params:      params,
			Middlewares: append(append([]Middleware{}, r.middlewares...), middlewares...),
		})
	}
}
