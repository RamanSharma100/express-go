package http

type Router struct {
	routes      []Route
	middlewares []Middleware
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) AddRoute(path string, handler Handler, method []string) {
	if validateRoute(path, handler) {

		if len(method) == 0 {
			method = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"}
		}

		params := []string{}

		if isParameterizedRoute(path) {
			path, params = getParameterizedRoute(path)
		}

		searchParams := getSearchParams(path)
		path = removeQueryParams(path)

		r.routes = append(r.routes, Route{
			Method:       method,
			Path:         path,
			Handler:      handler,
			Params:       params,
			SearchParams: searchParams,
			Middlewares:  append([]Middleware{}, r.middlewares...),
		})
	}
}

func (r *Router) Get(path string, handler Handler) *RouteChain {
	r.AddRoute(path, handler, []string{"GET"})
	return &RouteChain{
		router: r,
		path:   path,
		method: []string{"GET"},
	}
}

func (r *Router) Post(path string, handler Handler) *RouteChain {
	r.AddRoute(path, handler, []string{"POST"})
	return &RouteChain{
		router: r,
		path:   path,
		method: []string{"POST"},
	}
}

func (r *Router) Put(path string, handler Handler) *RouteChain {
	r.AddRoute(path, handler, []string{"PUT"})
	return &RouteChain{
		router: r,
		path:   path,
		method: []string{"PUT"},
	}
}

func (r *Router) Delete(path string, handler Handler) *RouteChain {
	r.AddRoute(path, handler, []string{"DELETE"})
	return &RouteChain{
		router: r,
		path:   path,
		method: []string{"DELETE"},
	}
}

func (r *Router) Patch(path string, handler Handler) *RouteChain {
	r.AddRoute(path, handler, []string{"PATCH"})
	return &RouteChain{
		router: r,
		path:   path,
		method: []string{"PATCH"},
	}
}

func (r *Router) Options(path string, handler Handler) *RouteChain {
	r.AddRoute(path, handler, []string{"OPTIONS"})
	return &RouteChain{
		router: r,
		path:   path,
		method: []string{"OPTIONS"},
	}
}

func (r *Router) Head(path string, handler Handler) *RouteChain {
	r.AddRoute(path, handler, []string{"HEAD"})
	return &RouteChain{
		router: r,
		path:   path,
		method: []string{"HEAD"},
	}
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
	if validateRoute(path, handler) {

		if len(method) == 0 {
			method = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"}
		}

		params := []string{}

		if isParameterizedRoute(path) {
			path, params = getParameterizedRoute(path)
		}

		searchParams := getSearchParams(path)
		path = removeQueryParams(path)

		r.routes = append(r.routes, Route{
			Method:       method,
			Path:         path,
			Handler:      handler,
			Params:       params,
			SearchParams: searchParams,
			Middlewares:  append(append([]Middleware{}, r.middlewares...), middlewares...),
		})
	}
}
