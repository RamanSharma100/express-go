package http

import "strings"

type Router struct {
	Routes []Route
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

		r.Routes = append(r.Routes, Route{
			Method:  method,
			Path:    path,
			Handler: handler,
			Params:  params,
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
